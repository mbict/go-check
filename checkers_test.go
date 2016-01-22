package check

import (
	. "gopkg.in/check.v1"
	"testing"
)

type CheckerSuite struct{}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&CheckerSuite{})

func (s *CheckerSuite) TestImplements(c *C) {
	var t Checker
	c.Assert(HasKey, Implements, &t)
	c.Assert(Each(nil), Implements, &t)
	c.Assert(Any(nil), Implements, &t)
}

func checkTest(c *C, checker Checker, obtained, expected interface{}, valid bool, message string) {
	res, err := checker.Check([]interface{}{obtained, expected}, []string{"", ""})

	c.Check(res, Equals, valid)
	c.Check(err, Equals, message)
}

func (s *CheckerSuite) TestHasKey(c *C) {
	obtained := map[string]interface{}{
		"id":  1234,
		"foo": "abc",
		"bar": nil,
	}

	checkTest(c, HasKey, nil, nil, false, "obtained should be of a map type map[string]interface{}")
	checkTest(c, HasKey, int64(0), nil, false, "obtained should be of a map type map[string]interface{}")
	checkTest(c, HasKey, obtained, nil, false, "expected keys should be of type []string")
	checkTest(c, HasKey, obtained, "test", false, "expected keys should be of type []string")
	checkTest(c, HasKey, obtained, []string{"id", "bar"}, true, "")
	checkTest(c, HasKey, obtained, []string{}, true, "")
	checkTest(c, HasKey, obtained, []string{"woes", "bar"}, false, "missing the following keys: woes")
}

func (s *CheckerSuite) TestEach(c *C) {
	checkTest(c, Each(Equals), nil, "foo", false, "input is not a slice or a array")
	checkTest(c, Each(Equals), int64(0), "foo", false, "input is not a slice or a array")
	checkTest(c, Each(Equals), []string{"foo", "foo"}, "foo", true, "")
	checkTest(c, Each(Equals), []string{}, "foo", true, "")
	checkTest(c, Each(Equals), []string{"bar"}, "foo", false, "")
	checkTest(c, Each(Equals), []string{"foo", "bar"}, "foo", false, "")
}

func (s *CheckerSuite) TestAny(c *C) {
	checkTest(c, Any(Equals), nil, "foo", false, "input is not a slice or a array")
	checkTest(c, Any(Equals), int64(0), "foo", false, "input is not a slice or a array")
	checkTest(c, Any(Equals), []string{"foo", "bar"}, "foo", true, "")
	checkTest(c, Any(Equals), []string{}, "foo", false, "")
	checkTest(c, Any(Equals), []string{"bar", "foo"}, "stub", false, "")
}
