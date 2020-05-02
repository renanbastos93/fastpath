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
It was compare method used currently on Fiber and origin URLPath, was tested on Windows and MacOS.

### Windows
```
goos: windows
goarch: amd64
pkg: github.com/renanbastos93/fastpath
BenchmarkRegexp-6        1879828               617 ns/op             304 B/op          3 allocs/op
BenchmarkUrlPath-6       7296402               160 ns/op             336 B/op          2 allocs/op
BenchmarkMatch-6         7871742               152 ns/op             336 B/op          2 allocs/op
PASS
ok      github.com/renanbastos93/fastpath       4.688s
``` 

### MacOS
```
goos: darwin
goarch: amd64
pkg: github.com/renanbastos93/fastpath
BenchmarkRegexp-4    	  962472	      1190 ns/op	     304 B/op	       3 allocs/op
BenchmarkUrlPath-4   	 4239836	       270 ns/op	     336 B/op	       2 allocs/op
BenchmarkMatch-4     	 4650106	       260 ns/op	     336 B/op	       2 allocs/op
PASS
ok  	github.com/renanbastos93/fastpath	5.116s
```
