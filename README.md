[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/set)](https://pkg.go.dev/github.com/s0rg/set)
[![License](https://img.shields.io/github/license/s0rg/set)](https://github.com/s0rg/set/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/set)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/set?sort=semver)](https://github.com/s0rg/set/tags)

[![CI](https://github.com/s0rg/set/workflows/ci/badge.svg)](https://github.com/s0rg/set/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/set)](https://goreportcard.com/report/github.com/s0rg/set)
[![Maintainability](https://api.codeclimate.com/v1/badges/aadc34c86aed23a42013/maintainability)](https://codeclimate.com/github/s0rg/set/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/aadc34c86aed23a42013/test_coverage)](https://codeclimate.com/github/s0rg/set/test_coverage)
![Issues](https://img.shields.io/github/issues/s0rg/set)


# set

generic set types for golang

# features

- both un-ordered and ordered types
- union, diff, intersect, comparision for any two sets
- simple API
- 100% test coverage

# example

```go
import (
    "fmt"

    "github.com/s0rg/set"
)

func main() {
    // create new, empty set of int's
    s := make(set.Unordered[int]) // fastest variant as it direct map functions call

    // or

    // second (slower) form for unordered constructor, it uses indirect calls via interface
    // s := set.NewUnordered[int]()

    // ordered constructor, only this form
    // s := set.NewOrdered[int]()

    // add some values
    s.Add(1)
    s.Add(2)

    // and some more...
    set.Load(s, 2, 3)

    // check set for value
    if !s.Has(2) {
        panic("2 not found")
    }

    // check and add
    if s.TryAdd(4) {
        fmt.Println("value 4 wasnt in set, it there now")
    }

    // delete item
    s.Del(1)

    fmt.Println("Set length:", s.Len())
    fmt.Println("Set contents:", s.ToSlice())

    s.Clear()

    fmt.Println("Set length:", s.Len())
    fmt.Println("Set contents:", s.ToSlice())
}
```
