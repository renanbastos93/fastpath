package fastpath

import (
	"strings"
)

type seg struct {
	Param       string
	Const       string
	ConstLength int
	IsParam     bool
	IsOptional  bool
}

// Path ...
type Path struct {
	Segs       []seg
	SegsLength int
	Params     []string
}

// New ...
func New(pattern string) (p Path) {
	aPattern := strings.Split(pattern, "/")[1:] // every route starts with an "/"
	var out = make([]seg, len(aPattern))
	var params []string
	for i := 0; i < len(aPattern); i++ {
		patternLen := len(aPattern[i])
		if patternLen == 0 { // skip empty parts
			continue
		}
		// is parameter
		if aPattern[i][0] == '*' || aPattern[i][0] == ':' {
			out[i] = seg{
				Param:      paramTrimmer(aPattern[i]),
				IsParam:    true,
				IsOptional: aPattern[i] == "*" || aPattern[i][patternLen-1] == '?',
			}
			params = append(params, out[i].Param)
		} else {
			out[i] = seg{
				Const: aPattern[i],
			}
		}
	}

	p = Path{Segs: out, SegsLength: len(out), Params: params}
	return
}

// Match ...
func (p *Path) Match(s string) ([]string, bool) {
	if len(s) > 0 {
		s = s[1:]
	}
	params := make([]string, len(p.Params), cap(p.Params))
	paramsIterator := 0
	for _, segment := range p.Segs {
		i := strings.IndexByte(s, '/')
		j := i + 1

		if i == -1 || (segment.IsParam && segment.Param == "*") {
			i = len(s)
			j = i
		}
		if segment.IsParam {
			if s[:i] == "" && !segment.IsOptional {
				return nil, false
			}
			params[paramsIterator] = s[:i]
			paramsIterator++
		} else {
			if s[:i] != segment.Const {
				return nil, false
			}
		}

		s = s[j:]
	}

	return params, true
}

func paramTrimmer(param string) string {
	start := 0
	end := len(param)

	if param[start] != ':' { // is not a param
		return param
	}
	start++
	if param[end-1] == '?' { // is ?
		end--
	}

	return param[start:end]
}
