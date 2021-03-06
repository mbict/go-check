package check

import (
	checkv1 "gopkg.in/check.v1"
	"reflect"
)

// Each checker loops over a slice an applies the provided checker on each element.
//
// For example:
//
//     a := []string{"a", "a"}
//     c.Assert(a, Each(Equals), b)
//
func Each(checker checkv1.Checker) checkv1.Checker {
	return &eachChecker{checker}
}

type eachChecker struct {
	sub checkv1.Checker
}

func (checker *eachChecker) Info() *checkv1.CheckerInfo {
	info := *checker.sub.Info()
	info.Name = "Each(" + info.Name + ")"
	return &info
}

func (checker *eachChecker) Check(params []interface{}, names []string) (result bool, error string) {
	t := reflect.TypeOf(params[0])
	if t == nil || (t.Kind() != reflect.Slice && t.Kind() != reflect.Array) {
		return false, "input is not a slice or a array"
	}

	v := reflect.ValueOf(params[0])
	for i := 0; i < v.Len(); i++ {
		result, error = checker.sub.Check(append([]interface{}{v.Index(i).Interface()}, params[1:]...), names)
		if !result {
			return result, error
		}
	}
	return true, ""
}
