package check

import (
	. "gopkg.in/check.v1"
	"reflect"
)

// The Any checker loops over a slice an applies the provided checker on any element.
// If one of the elements in the slice matches the provided checker it will succeed otherwise fail
//
// For example:
//
//     a := []string{"a", "a"}
//     c.Assert(a, Any(Equals), b)
//
func Any(checker Checker) Checker {
	return &anyChecker{checker}
}

type anyChecker struct {
	sub Checker
}

func (checker *anyChecker) Info() *CheckerInfo {
	info := *checker.sub.Info()
	info.Name = "Any(" + info.Name + ")"
	return &info
}

func (checker *anyChecker) Check(params []interface{}, names []string) (bool, string) {
	t := reflect.TypeOf(params[0])
	if t == nil || (t.Kind() != reflect.Slice && t.Kind() != reflect.Array) {
		return false, "input is not a slice or a array"
	}

	v := reflect.ValueOf(params[0])
	for i := 0; i < v.Len(); i++ {
		result, _ := checker.sub.Check(append([]interface{}{v.Index(i).Interface()}, params[1:]...), names)
		if result {
			return true, ""
		}
	}
	return false, ""
}
