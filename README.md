![Go](https://github.com/renanbastos93/fastpath/workflows/Go/badge.svg)

# fastpath
This lib was based on [URLPATH](https://github.com/ucarion/urlpath) created by [@ucarion](https://github.com/ucarion), the start was made a fork it. Even so, we created a new repository without the fork. But the credits are entirely from Ucarion.

# How was born it?
We need to get parameters optional to use in [Fiber](https://gofiber.io/) because we have a big need to improve the router. Until the present moment, it uses regex to validate then we go to remove with this lib.

# How to use...
```go 
package main

import (
    "fmt"

    "github.com/renanbastos93/fastpath"
)

func main() {
	p := fastpath.New("/api/user/:id")
    params, ok := p.Match("/api/user/728342")

    if !ok {
        // not match
        return
    }
    // Matched and have parameters, so will return a map
    fmt.Println(params["id"]) // 728342
}
```

# Use cases
It was created some use cases to validate this approach to use on Fiber. Note: Wildcard and parameter optional only can use on the last path. You can see more examples on unit tests.

# Performance
It was compare method used currently on Fiber and origin URLPath, was tested on MacOS.
### MacOS
```
goos: darwin
goarch: amd64
pkg: github.com/renanbastos93/fastpath
BenchmarkMatch-12                                                           21682538   55.4  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//_-_/-12                                                75846546   15.5  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/*_-_/api/v1/entity/1-12                            25320554   47.5  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/*_-_/api/v1/entity/1-12                  13227706   90.0  ns/op  32   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2/:nomatch_-_/api/v1/entity/1-12   12470054   96.4  ns/op  64   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2/:param3_-_/api/v1/entity/1/2-12  11824402   101   ns/op  64   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2?_-_/api/v1/entity/1-12           14237450   85.2  ns/op  32   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param2?_-_/api/v1/-12                          22081225   54.7  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param?_-_/api/v1/entity-12                     21206205   54.9  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/const_-_/api/v1/const-12                        38504092   31.9  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/test_-_/api/v1/noMatch-12                       38351829   32.2  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases/_-_/-12                                                 149556266  8.08  ns/op  0    B/op  0  allocs/op
BenchmarkRegexp-12                                                          1799762    662   ns/op  304  B/op  3  allocs/op
BenchmarkUrlPath-12                                                         7292202    166   ns/op  336  B/op  2  allocs/op
```
