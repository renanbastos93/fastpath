package fastpath

import (
	"strings"
)

type seg struct {
	Param      string
	Const      string
	IsParam    bool
	IsOptional bool
	IsWildcard bool
}

// Path ...
type Path struct {
	S []seg
}

// New ...
func New(pattern string) (p Path) {
	if pattern == "*" {
		return
	}
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
					Param:      aPattern[i],
					IsParam:    true,
					IsOptional: true,
				}
			} else {
				out[i] = seg{
					Param:   aPattern[i],
					IsParam: true,
				}
			}
		} else if aPattern[i] == "*" {
			hasOpt = true
			out[i] = seg{
				Param:      aPattern[i],
				IsParam:    true,
				IsWildcard: true,
			}
		} else {
			out[i] = seg{
				Const: aPattern[i],
			}
		}
	}
	p = Path{S: out[1:]}
	return
}

// Match ...
func (p *Path) Match(s string) (map[string]string, bool) {
	params := map[string]string{}
	/*
		DEGRADED PERFORMANCE
		SO WE COMMENTED THIS BLOCK
		if s[0:1] == "/" && s[len(s)-1:] != "/" {
			s = s[1:] + "/"
		} else {
			s = s[1:]
		}
	*/
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
			if segment.IsOptional {
				params[segment.Param[1:len(segment.Param)-1]] = s[:i]
				continue
			}
			if segment.IsWildcard {
				params[segment.Param] = s[:i]
				continue
			}
			if s[:i] == "" {
				return nil, false
			}
			params[segment.Param[1:]] = s[:i]
		} else {
			if s[:i] != segment.Const {
				return nil, false
			}
		}

		s = s[j:]
	}
	return params, true
}
