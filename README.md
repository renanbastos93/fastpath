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
BenchmarkMatch-12                                                           69732186   16.3  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//_-_/-12                                                100000000  10.7  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/*_-_/api/v1/entity/1-12                            75475124   17.2  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/*/:param3_-_/api/v1/entity/1/2-12               32059923   38.8  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/:param/*_-_/api/v1/entity/1-12                  42632986   25.0  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2/:nomatch_-_/api/v1/entity/1-12   38993047   32.0  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2/:param3_-_/api/v1/entity/1/2-12  35778060   33.8  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/:param/:param2?_-_/api/v1/entity/1-12           48683749   25.2  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/:param2?_-_/api/v1/-12                          79739908   15.2  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/:param?_-_/api/v1/entity-12                     77837749   15.8  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/const_-_/api/v1/const-12                        100000000  10.1  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//api/v1/test_-_/api/v1/noMatch-12                       138586831  8.59  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//config/abc.json_-_/config/abc.json-12                  100000000  10.7  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases//config/noMatch.json_-_/config/abc.json-12              185185951  6.57  ns/op  0    B/op  0  allocs/op
BenchmarkMatchCases/_-_/-12                                                 100000000  10.0  ns/op  0    B/op  0  allocs/op
BenchmarkRegexp-12                                                          1792290    675   ns/op  304  B/op  3  allocs/op
BenchmarkUrlPath-12                                                         7052295    217   ns/op  336  B/op  2  allocs/op
```
