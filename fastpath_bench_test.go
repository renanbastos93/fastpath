package fastpath

import (
	"regexp"
	"strings"
	"testing"

	"github.com/ucarion/urlpath"
)

const pattern = "/api/user/:id"
const uri = "/api/user/728342"

func getRegex(p string) (*regexp.Regexp, error) {
	pattern := "^"
	segments := strings.Split(p, "/")
	for _, s := range segments {
		if s == "" {
			continue
		}
		if s[0] == ':' {
			if strings.Contains(s, "?") {
				pattern += "(?:/([^/]+?))?"
			} else {
				pattern += "/(?:([^/]+?))"
			}
		} else if s[0] == '*' {
			pattern += "/(.*)"
		} else {
			pattern += "/" + s
		}
	}
	pattern += "/?$"
	regex, err := regexp.Compile(pattern)
	return regex, err
}

func matchRegex(regex *regexp.Regexp, p string) (match bool, values []string) {
	if regex.MatchString(p) {
		// get values for parameters
		matches := regex.FindAllStringSubmatch(p, -1)
		// did we get the values?
		if len(matches) > 0 && len(matches[0]) > 1 {
			values = matches[0][1:len(matches[0])]
			return true, values
		}
		return false, values
	}
	return false, values
}

func BenchmarkRegexp(b *testing.B) {
	regex, _ := getRegex(pattern)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = matchRegex(regex, uri)
	}
}

func BenchmarkUrlPath(b *testing.B) {
	parser := urlpath.New(pattern)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = parser.Match(uri)
	}
}

func BenchmarkMatch(b *testing.B) {
	p := New(pattern)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = p.Match(uri)
	}
}
