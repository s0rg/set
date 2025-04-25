[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/set)](https://pkg.go.dev/github.com/s0rg/set)
[![License](https://img.shields.io/github/license/s0rg/set)](https://github.com/s0rg/set/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/set)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/set?sort=semver)](https://github.com/s0rg/set/tags)

[![CI](https://github.com/s0rg/set/workflows/ci/badge.svg)](https://github.com/s0rg/set/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/set)](https://goreportcard.com/report/github.com/s0rg/set)
[![Maintainability](https://qlty.sh/badges/06d50a96-bbc8-4635-9f48-83ea9fdf548b/maintainability.svg)](https://qlty.sh/gh/s0rg/projects/set)
[![Code Coverage](https://qlty.sh/badges/06d50a96-bbc8-4635-9f48-83ea9fdf548b/test_coverage.svg)](https://qlty.sh/gh/s0rg/projects/set)
![Issues](https://img.shields.io/github/issues/s0rg/set)

# set

generic set types for golang

# features

- both un-ordered and ordered types
- union, diff, intersect, comparision for any two sets
- simple API
- zero-alloc
- zero-dependency
- 100% test coverage

# example

```go
import (
    "fmt"

    "github.com/s0rg/set"
)

func main() {
    // create new, empty set of int's
    s := make(set.Unordered[int]) // fastest variant as it direct functions call

    // or

    // second (a bit slower) form for unordered constructor, it uses indirect calls via interface
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

    // iter over items
    for i := range s.Iter {
        fmt.Printf("iter: %d\n", i)
    }

    s.Clear()

    fmt.Println("Set length:", s.Len())
    fmt.Println("Set contents:", s.ToSlice())
}
```

# benchmarks
```
cpu: AMD Ryzen 5 5500U with Radeon Graphics
BenchmarkSetUnorderedDirect/Add-12      7092192        174.1 ns/op      49 B/op          0 allocs/op
BenchmarkSetUnorderedDirect/Has-12      17680040       97.33 ns/op       0 B/op          0 allocs/op
BenchmarkSetUnorderedDirect/Len-12      1000000000    0.2496 ns/op       0 B/op          0 allocs/op
BenchmarkSetUnorderedDirect/Pop-12      6251703         5282 ns/op       0 B/op          0 allocs/op
BenchmarkSetUnorderedDirect/Del-12      561568860      2.145 ns/op       0 B/op          0 allocs/op
BenchmarkSetUnorderedIndirect/Add-12    7479926        193.6 ns/op      46 B/op          0 allocs/op
BenchmarkSetUnorderedIndirect/Has-12    17168162       80.38 ns/op       0 B/op          0 allocs/op
BenchmarkSetUnorderedIndirect/Len-12    802026860      1.496 ns/op       0 B/op          0 allocs/op
BenchmarkSetUnorderedIndirect/Pop-12    6782689         6487 ns/op       0 B/op          0 allocs/op
BenchmarkSetUnorderedIndirect/Del-12    393659181      3.072 ns/op       0 B/op          0 allocs/op
BenchmarkSetOrdered/Add-12              6576442        178.7 ns/op      67 B/op          0 allocs/op
BenchmarkSetOrdered/Has-12              17831434       84.19 ns/op       0 B/op          0 allocs/op
BenchmarkSetOrdered/Len-12              473897580      2.494 ns/op       0 B/op          0 allocs/op
BenchmarkSetOrdered/Pop-12              297457706      4.054 ns/op       0 B/op          0 allocs/op
BenchmarkSetOrdered/Del-12              144014604      8.326 ns/op       0 B/op          0 allocs/op
```
