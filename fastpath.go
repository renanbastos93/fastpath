package fastpath

import (
	"strings"
)

type seg struct {
	Param      string
	Const      string
	IsParam    bool
	IsOptional bool
}

// Path ...
type Path struct {
	S []seg
}

// New ...
func New(pattern string) (p Path) {
	aPattern := strings.Split(pattern, "/")
	var hasOpt bool = false
	var out = make([]seg, len(aPattern))
	for i := 0; i < len(aPattern); i++ {
		if hasOpt && i < len(aPattern) {
			panic("malformed pattern")
		}
		if strings.HasPrefix(aPattern[i], ":") {
			if strings.HasSuffix(aPattern[i], "?") {
				hasOpt = true
				out[i] = seg{
					Param:      aPattern[i][1 : len(aPattern[i])-1],
					IsParam:    true,
					IsOptional: true,
				}
			} else {
				out[i] = seg{
					Param:   aPattern[i][1:],
					IsParam: true,
				}
			}
		} else if aPattern[i] == "*" {
			hasOpt = true
			out[i] = seg{
				Param:      aPattern[i],
				IsParam:    true,
				IsOptional: true,
			}
		} else {
			out[i] = seg{
				Const: aPattern[i],
			}
		}
	}
	p = Path{S: out}
	return
}

// Match ...
func (p *Path) Match(s string) (map[string]string, bool) {
	params := map[string]string{}
	for segmentIndex, segment := range p.S {
		i := strings.IndexByte(s, '/')
		j := i + 1

		if i == -1 {
			i = len(s)
			j = len(s)
			if segmentIndex != len(p.S)-1 {
				return nil, false
			}
		} else {
			if segmentIndex == len(p.S)-1 {
				return nil, false
			}
		}
		if segment.IsParam {
			if s[:i] == "" && !segment.IsOptional {
				return nil, false
			}
			params[segment.Param] = s[:i]
		} else {
			if s[:i] != segment.Const {
				return nil, false
			}
		}

		s = s[j:]

	}
	return params, true
}
