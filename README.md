[![wercker status](https://app.wercker.com/status/b7b9694a4d1fd415c6d666ba4cf4d01a/s "wercker status")](https://app.wercker.com/project/bykey/b7b9694a4d1fd415c6d666ba4cf4d01a)
[![Build Status](https://travis-ci.org/mbict/go-check.png?branch=master)](https://travis-ci.org/mbict/go-check)
[![GoDoc](https://godoc.org/github.com/mbict/go-check?status.png)](http://godoc.org/github.com/mbict/go-check)
[![GoCover](http://gocover.io/_badge/github.com/mbict/go-check)](http://gocover.io/github.com/mbict/go-check)

Check
=====

Check adds new checkers for the [labix check.v1](https://github.com/go-check/check/tree/v1) testing library

HasKey
======
HasKey checks if a key exists in a ```map[string]interface{}``` typed map

#### example:
```go
obtained := map[string]interface{}{
    "id":  1234,
    "foo": "abc",
    "bar": nil,
}

c.Assert(obtained, HasKey, []string{"foo", "bar"})
```

Each
====
Each iterates over a slice/array and runs the embedded checker for each record
All record must match to get a positive check result.

#### example with the HasKey chained:
```go
obtained := []map[string]interface{}{
    {
        "id":  1234,
        "foo": "abc",
        "bar": nil,
    }, {
        "id":  5678,
        "foo": "xyz",
        "bar": "aaa",
    },
}
c.Assert(obtained, Each(HasKey), []string{"foo", "bar"})
```

Any
====
Any iterates over a slice/array and runs the embedded checker for each record.
One or more records must match to get a positive match.

#### example with the HasKey chained:
```go
obtained := []string{"bar", "foo"}
c.Assert(obtained, Any(Equals), "foo")
```
