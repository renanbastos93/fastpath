package main

import (
	"fmt"
	"path"
	"regexp"
	"testing"
)

// path.Match("/api/a/?", "/api/a/asdada")
// path.Match("/api/?/b", "/api/asda/b")
// path.Match("/api/a/b/*", "/api/a/b/asduahuda")
// path.Match("/api/c/b", "/api/c/b")

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New("/api/v1/:param/*")
	}
}

func BenchmarkMatchNative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		path.Match("/api/v1/?/*", "/api/v1/entity/id")
	}
}

func BenchmarkMatch(b *testing.B) {
	p := New("/api/v1/:param/*")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		p.Match("/api/v1/entity/id")
	}
}

func BenchmarkRegex(b *testing.B) {
	regex := regexp.MustCompile("/api/v1/(?P<param>[^/]+)/.*")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		regex.FindStringSubmatch(fmt.Sprintf("/api/v1/param%d/extra", i))
	}
}

// go test -benchmem -run=^$ -bench . | /Users/renanbastos/go/bin/prettybench
// go test -benchmem -run=^$ -bench .
// goos: darwin
// goarch: amd64
// pkg: github.com/renanbastos93/fastpath
// BenchmarkNew-4     	 3344361	       350 ns/op	     288 B/op	       2 allocs/op
// BenchmarkMatch-4   	 2772646	       424 ns/op	     416 B/op	       3 allocs/op
// BenchmarkRegex-4   	 1340355	       903 ns/op	      72 B/op	       3 allocs/op
// PASS
// ok  	github.com/renanbastos93/fastpath	5.273s
