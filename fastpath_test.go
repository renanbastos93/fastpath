package fastpath

import (
	"reflect"
	"testing"
)

func TestPath_Match(t *testing.T) {
	type cases struct {
		uri    string
		params map[string]string
		ok     bool
	}
	tests := []struct {
		name    string
		pattern Path
		cases   []cases
	}{
		// Pattern: /api/v1/:param/*
		{
			name:    "For match URL to Pattern with a param and wildcard",
			pattern: New("/api/v1/:param/*"),
			cases: []cases{
				{
					uri:    "/api/v1/entity",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v1/entity/",
					params: map[string]string{"param": "entity", "*": ""},
					ok:     true,
				},
				{
					uri:    "/api/v1/entity/1",
					params: map[string]string{"param": "entity", "*": "1"},
					ok:     true,
				},
				{
					uri:    "/api/v",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v2",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v1/",
					params: nil,
					ok:     false,
				},
			},
		},
		// Pattern: /api/v1/:param?
		{
			name:    "For match URL to Pattern with a param optional",
			pattern: New("/api/v1/:param?"),
			cases: []cases{
				{
					uri:    "/api/v1",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v1/",
					params: map[string]string{"param": ""},
					ok:     true,
				},
				{
					uri:    "/api/v1/optional",
					params: map[string]string{"param": "optional"},
					ok:     true,
				},
				{
					uri:    "/api/v",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v2",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/xyz",
					params: nil,
					ok:     false,
				},
			},
		},
		// Pattern: /api/v1/*
		{
			name:    "For match URL to Pattern with wildcard",
			pattern: New("/api/v1/*"),
			cases: []cases{
				{
					uri:    "/api/v1",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v1/",
					params: map[string]string{"*": ""},
					ok:     true,
				},
				{
					uri:    "/api/v1/entity",
					params: map[string]string{"*": "entity"},
					ok:     true,
				},
				{
					uri:    "/api/v",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v2",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/abc",
					params: nil,
					ok:     false,
				},
			},
		},
		// Pattern: /api/v1/:param
		{
			name:    "For match URL to Pattern with a param",
			pattern: New("/api/v1/:param"),
			cases: []cases{
				{
					uri:    "/api/v1/entity",
					params: map[string]string{"param": "entity"},
					ok:     true,
				},
				{
					uri:    "/api/v1",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v1/",
					params: nil,
					ok:     false,
				},
			},
		},
		// Pattern: /api/v1/const
		{
			name:    "For match URL to Pattern without a param or wildcard",
			pattern: New("/api/v1/const"),
			cases: []cases{
				{
					uri:    "/api/v1/const",
					params: map[string]string{},
					ok:     true,
				},
				{
					uri:    "/api/v1",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v1/",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/api/v1/something",
					params: nil,
					ok:     false,
				},
			},
		},
		// Pattern: /api/v1/:param/abc/*
		{
			name:    "For match URL to Pattern with a param and wildcard differents position",
			pattern: New("/api/v1/:param/abc/*"),
			cases: []cases{
				{
					uri:    "/api/v1/well/abc/wildcard",
					params: map[string]string{"param": "well", "*": "wildcard"},
					ok:     true,
				},
				{
					uri:    "/api/v1/well/abc/",
					params: map[string]string{"param": "well", "*": ""},
					ok:     true,
				},
				{
					uri:    "/api/v1/well/abc",
					params: nil,
					ok:     false,
				},
			},
		},
		// Pattern: /:day?/:month?/:year?
		{
			name:    "For match URL to Pattern with paremeters optional",
			pattern: New("/api/:day/:month?/:year?"),
			cases: []cases{
				{
					uri:    "/api/1",
					params: nil,
					ok:     false,
				}, {
					uri:    "/api/1/",
					params: map[string]string{"day": "1", "month": "", "year": ""},
					ok:     true,
				},
				{
					uri:    "/api/1/2",
					params: map[string]string{"day": "1", "month": "2", "year": ""},
					ok:     true,
				},
				{
					uri:    "/api/1/2/3",
					params: map[string]string{"day": "1", "month": "2", "year": "3"},
					ok:     true,
				},
			},
		},
		// Pattern: /api/*
		{
			name:    "Wildcard simple",
			pattern: New("/api/*"),
			cases: []cases{
				{
					uri:    "/api/",
					params: map[string]string{"*": ""},
					ok:     true,
				},
				{
					uri:    "/api/joker",
					params: map[string]string{"*": "joker"},
					ok:     true,
				},
				{
					uri:    "/api",
					params: nil,
					ok:     false,
				},
			},
		},
		// Pattern: /
		{
			name:    "Simple path",
			pattern: New("/"),
			cases: []cases{
				{
					uri:    "/api",
					params: nil,
					ok:     false,
				},
				{
					uri:    "",
					params: nil,
					ok:     false,
				},
				{
					uri:    "/",
					params: map[string]string{},
					ok:     true,
				},
			},
		},
		// with not match with simple pattern
		{
			name:    "Simple path",
			pattern: New("/xyz"),
			cases: []cases{
				{
					uri:    "xyz",
					params: nil,
					ok:     false,
				},
				{
					uri:    "xyz/",
					params: nil,
					ok:     false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tcase := range tt.cases {
				params, ok := tt.pattern.Match(tcase.uri)
				if !reflect.DeepEqual(params, tcase.params) {
					t.Errorf("Path.Match() got = %v, want %v", params, tcase.params)
				}
				if ok != tcase.ok {
					t.Errorf("Path.Match() got1 = %v, want %v", ok, tcase.ok)
				}
			}
		})
	}
}

// go test -coverprofile "coverage.html" "github.com/renanbastos93/fastpath" . && go tool cover -func="coverage.html"
// github.com/renanbastos93/fastpath/fastpath.go:20:       New             100.0%
// github.com/renanbastos93/fastpath/fastpath.go:60:       Match           94.7%
// total:                                                  (statements)    97.2%
