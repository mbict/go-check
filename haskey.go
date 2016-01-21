package check

import (
	"fmt"
	. "gopkg.in/check.v1"
	"strings"
)

type keyChecker struct {
	*CheckerInfo
}

var HasKey Checker = &keyChecker{
	&CheckerInfo{Name: "HasKey", Params: []string{"obtained", "keys"}},
}

func (checker *keyChecker) Check(params []interface{}, names []string) (result bool, error string) {

	values, ok := params[0].(map[string]interface{})
	if !ok {
		return false, "obtained should be of a map type map[string]interface{}"
	}

	keys, ok := params[1].([]string)
	if !ok {
		return false, "expected keys should be of type []string"
	}

	defer func() {
		if v := recover(); v != nil {
			result = false
			error = fmt.Sprint(v)
		}
	}()

	missingKeys := []string{}
	for _, key := range keys {
		if _, ok := values[key]; !ok {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		return false, "missing the following keys: " + strings.Join(missingKeys, ", ")
	}

	return true, ""
}
