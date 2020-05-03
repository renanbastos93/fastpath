package fastpath

import (
	"fmt"
	"reflect"
	"testing"
)

type testcase struct {
	uri    string
	params []string
	ok     bool
}

func Test_With_Param_And_Wildcard(t *testing.T) {
	checkCases(
		t,
		New("/api/v1/:param/*"),
		[]testcase{
			{uri: "/api/v1/entity", params: []string{"entity", ""}, ok: true},
			{uri: "/api/v1/entity/", params: []string{"entity", ""}, ok: true},
			{uri: "/api/v1/entity/1", params: []string{"entity", "1"}, ok: true},
			{uri: "/api/v", params: nil, ok: false},
			{uri: "/api/v2", params: nil, ok: false},
			{uri: "/api/v1/", params: nil, ok: false},
		},
	)
}

func Test_With_A_Param_Optional(t *testing.T) {
	checkCases(
		t,
		New("/api/v1/:param?"),
		[]testcase{
			{uri: "/api/v1", params: []string{""}, ok: true},
			{uri: "/api/v1/", params: []string{""}, ok: true},
			{uri: "/api/v1/optional", params: []string{"optional"}, ok: true},
			{uri: "/api/v", params: nil, ok: false},
			{uri: "/api/v2", params: nil, ok: false},
			{uri: "/api/xyz", params: nil, ok: false},
		},
	)
}

func Test_With_With_Wildcard(t *testing.T) {
	checkCases(
		t,
		New("/api/v1/*"),
		[]testcase{
			{uri: "/api/v1", params: []string{""}, ok: true},
			{uri: "/api/v1/", params: []string{""}, ok: true},
			{uri: "/api/v1/entity", params: []string{"entity"}, ok: true},
			{uri: "/api/v1/entity/1/2", params: []string{"entity/1/2"}, ok: true},
			{uri: "/api/v", params: nil, ok: false},
			{uri: "/api/v2", params: nil, ok: false},
			{uri: "/api/abc", params: nil, ok: false},
		},
	)
}
func Test_With_With_Param(t *testing.T) {
	checkCases(
		t,
		New("/api/v1/:param"),
		[]testcase{
			{uri: "/api/v1/entity", params: []string{"entity"}, ok: true},
			{uri: "/api/v1", params: nil, ok: false},
			{uri: "/api/v1/", params: nil, ok: false},
		},
	)
}
func Test_With_Without_A_Param_Or_Wildcard(t *testing.T) {
	checkCases(
		t,
		New("/api/v1/const"),
		[]testcase{
			{uri: "/api/v1/const", params: []string{}, ok: true},
			{uri: "/api/v1", params: nil, ok: false},
			{uri: "/api/v1/", params: nil, ok: false},
			{uri: "/api/v1/something", params: nil, ok: false},
		},
	)
}
func Test_With_With_A_Param_And_Wildcard_Differents_Positions(t *testing.T) {
	checkCases(
		t,
		New("/api/v1/:param/abc/*"),
		[]testcase{
			{uri: "/api/v1/well/abc/wildcard", params: []string{"well", "wildcard"}, ok: true},
			{uri: "/api/v1/well/abc/", params: []string{"well", ""}, ok: true},
			{uri: "/api/v1/well/abc", params: []string{"well", ""}, ok: true},
			{uri: "/api/v1/well/ttt", params: nil, ok: false},
		},
	)
}
func Test_With_With_Params_And_Optional(t *testing.T) {
	checkCases(
		t,
		New("/api/:day/:month?/:year?"),
		[]testcase{
			{uri: "/api/1", params: []string{"1", "", ""}, ok: true},
			{uri: "/api/1/", params: []string{"1", "", ""}, ok: true},
			{uri: "/api/1/2", params: []string{"1", "2", ""}, ok: true},
			{uri: "/api/1/2/3", params: []string{"1", "2", "3"}, ok: true},
			{uri: "/api/", params: nil, ok: false},
		},
	)
}
func Test_With_With_Simple_Wildcard(t *testing.T) {
	checkCases(
		t,
		New("/api/*"),
		[]testcase{
			{uri: "/api/", params: []string{""}, ok: true},
			{uri: "/api/joker", params: []string{"joker"}, ok: true},
			{uri: "/api", params: []string{""}, ok: true},
		},
	)
}
func Test_With_With_Wildcard_And_Optional(t *testing.T) {
	checkCases(
		t,
		New("/api/*/:param?"),
		[]testcase{
			{uri: "/api/", params: []string{"", ""}, ok: true},
			{uri: "/api/joker", params: []string{"joker", ""}, ok: true},
			{uri: "/api/joker/batman", params: []string{"joker", "batman"}, ok: true},
			{uri: "/api/joker/batman/robin", params: []string{"joker/batman", "robin"}, ok: true},
			{uri: "/api/joker/batman/robin/1", params: []string{"joker/batman/robin", "1"}, ok: true},
			{uri: "/api", params: []string{"", ""}, ok: true},
		},
	)
}
func Test_With_With_Wildcard_And_Param(t *testing.T) {
	checkCases(
		t,
		New("/api/*/:param"),
		[]testcase{
			{uri: "/api/test/abc", params: []string{"test", "abc"}, ok: true},
			{uri: "/api/joker/batman", params: []string{"joker", "batman"}, ok: true},
			{uri: "/api/joker/batman/robin", params: []string{"joker/batman", "robin"}, ok: true},
			{uri: "/api/joker/batman/robin/1", params: []string{"joker/batman/robin", "1"}, ok: true},
			{uri: "/api", params: nil, ok: false},
		},
	)
}
func Test_With_With_Wildcard_And_2Params(t *testing.T) {
	checkCases(
		t,
		New("/api/*/:param/:param2"),
		[]testcase{
			{uri: "/api/test/abc", params: nil, ok: false},
			{uri: "/api/joker/batman", params: nil, ok: false},
			{uri: "/api/joker/batman/robin", params: []string{"joker", "batman", "robin"}, ok: true},
			{uri: "/api/joker/batman/robin/1", params: []string{"joker/batman", "robin", "1"}, ok: true},
			{uri: "/api/joker/batman/robin/1/2", params: []string{"joker/batman/robin", "1", "2"}, ok: true},
			{uri: "/api", params: nil, ok: false},
		},
	)
}
func Test_With_With_Simple_Path(t *testing.T) {
	checkCases(
		t,
		New("/"),
		[]testcase{
			{uri: "/api", params: nil, ok: false},
			{uri: "", params: []string{}, ok: true},
			{uri: "/", params: []string{}, ok: true},
		},
	)
}
func Test_With_With_Empty_Path(t *testing.T) {
	checkCases(
		t,
		New(""),
		[]testcase{
			{uri: "/api", params: nil, ok: false},
			{uri: "", params: []string{}, ok: true},
			{uri: "/", params: []string{}, ok: true},
		},
	)
}
func Test_With_With_Simple_Path_And_NoMatch(t *testing.T) {
	checkCases(
		t,
		New("/xyz"),
		[]testcase{
			{uri: "xyz", params: nil, ok: false},
			{uri: "xyz/", params: nil, ok: false},
		},
	)
}

func checkCases(tParent *testing.T, parser Path, tcases []testcase) {
	for _, tcase := range tcases {
		tParent.Run(fmt.Sprintf("%+v", tcase), func(t *testing.T) {
			params, ok := parser.Match(tcase.uri)
			if !reflect.DeepEqual(params, tcase.params) {
				t.Errorf("Path.Match() got = %v, want %v", params, tcase.params)
			}
			if ok != tcase.ok {
				t.Errorf("Path.Match() got1 = %v, want %v", ok, tcase.ok)
			}
		})
	}
}

// go test -coverprofile "coverage.html" "github.com/renanbastos93/fastpath" . && go tool cover -func="coverage.html"
// github.com/renanbastos93/fastpath/fastpath.go:20:       New             100.0%
// github.com/renanbastos93/fastpath/fastpath.go:60:       Match           94.7%
// total:                                                  (statements)    97.2%
