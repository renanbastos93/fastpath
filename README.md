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
BenchmarkMatch-12                                                           28915200  41.7  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//_-_/-12                                                89510649  13.9  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/*_-_/api/v1/entity/1-12                            28645147  41.8  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/*/:param3_-_/api/v1/entity/1/2-12               13452026  90.3  ns/op  32   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/*_-_/api/v1/entity/1-12                  15885230  74.4  ns/op  32   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2/:nomatch_-_/api/v1/entity/1-12   14090077  85.5  ns/op  64   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2/:param3_-_/api/v1/entity/1/2-12  13904775  87.1  ns/op  64   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2?_-_/api/v1/entity/1-12           16266535  74.4  ns/op  32   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param2?_-_/api/v1/-12                          27986340  42.2  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/:param?_-_/api/v1/entity-12                     27915159  42.9  ns/op  16   B/op  1  allocs/op
BenchmarkMatchCases//api/v1/const_-_/api/v1/const-12                        85630245  15.0  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/test_-_/api/v1/noMatch-12                       94954256  12.7  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases/_-_/-12                                                 88474910  13.7  ns/op  0    B/op  0  allocs/op
BenchmarkRegexp-12                                                          1807408   662   ns/op  304  B/op  3  allocs/op
BenchmarkUrlPath-12                                                         7256646   167   ns/op  336  B/op  2  allocs/op
```
