package fastpath

import (
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
		New("/api/v1/:param/abc/*")
	}
}

func BenchmarkMatchNative(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		path.Match("/api/?/abc/?", "/api/v1/vamo/abc/optional")
	}
}

func BenchmarkMatch(b *testing.B) {
	p := New("/api/v1/:param/abc/*")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Match("/api/v1/vamo/abc/optional")
	}
}

func BenchmarkRegex(b *testing.B) {
	regex := regexp.MustCompile("/api/v1/(?P<param>[^/]+)/abc/(.*)")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		regex.FindStringSubmatch("/api/v1/vamo/abc/optional")
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

// WINDOWS
// goos: windows
// goarch: amd64
// pkg: github.com/renanbastos93/fastpath
// BenchmarkNew-6                   4658312               251 ns/op             336 B/op          2 allocs/op
// BenchmarkMatchNative-6          35389251                33.1 ns/op             0 B/op          0 allocs/op
// BenchmarkMatch-6                28632032                41.2 ns/op            48 B/op          1 allocs/op
// BenchmarkRegex-6                 2697618               441 ns/op             112 B/op          2 allocs/op
// PASS
// ok      github.com/renanbastos93/fastpath       5.677s
