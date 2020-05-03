![Go](https://github.com/renanbastos93/fastpath/workflows/Go/badge.svg)

# fastpath
This library was based on [urlpath](https://github.com/ucarion/urlpath) created by [@ucarion](https://github.com/ucarion). It started as a fork, but we've eventually decided to rewrite it from the ground up, based on the original code. All credits for the original library go to @ucarion.

## How was it born?
We wanted to come up with a route matching strategy for [Fiber](https://gofiber.io/) because at the time this library was created, it used regex for this purpose. Go's regex is currently very slow compared to other languages, so in order to achieve the best performance, we had to do it our own way.

## Usage
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
    // Matched and have parameters, so will return a slice
    fmt.Println(params[0]) // 728342
}
```

## Performance
We have compared the performance of `fastpath` with Fiber's regex matcher and the original `urlpath` library.

TODO: update benchmarks

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
