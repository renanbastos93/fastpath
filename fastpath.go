package fastpath

import (
	"strings"
)

type seg struct {
	Param      string
	Const      string
	IsParam    bool
	IsOptional bool
	IsLast     bool
}

// Path ...
type Path struct {
	Segs   []seg
	Params []string
}

// New ...
func New(pattern string) (p Path) {
	aPattern := strings.Split(pattern, "/")[1:] // every route starts with an "/"
	patternLen := len(aPattern)
	var out = make([]seg, patternLen)
	var params []string
	for i := 0; i < patternLen; i++ {
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
	if patternLen != 0 {
		out[patternLen-1].IsLast = true
	}
	p = Path{Segs: out, Params: params}
	return
}

// Match ...
func (p *Path) Match(s string) ([]string, bool) {
	if len(s) > 0 {
		s = s[1:]
	}
	params := make([]string, len(p.Params), cap(p.Params))
	paramsIterator := 0
	for index, segment := range p.Segs {
		i := strings.IndexByte(s, '/')
		j := i + 1

		if i == -1 || (segment.IsLast && segment.IsParam && segment.Param == "*") {
			i = len(s)
			j = i
		} else if !segment.IsLast && segment.IsParam && segment.Param == "*" {
			// for the expressjs behavior -> "/api/*/:param" - "/api/joker/batman/robin/1" -> "joker/batman/robin", "1"
			i = findCharPos(s, '/', strings.Count(s, "/")-(len(p.Segs)-(index+1))+1)
			j = i + 1
		}
		if segment.IsParam {
			if !segment.IsOptional && s[:i] == "" {
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
func findCharPos(s string, char byte, matchCount int) int {
	if matchCount == 0 {
		matchCount = 1
	}
	endPos, pos := 0, 0
	for matchCount > 0 && pos != -1 {
		if pos > 0 {
			s = s[pos+1:]
			endPos++
		}
		pos = strings.IndexByte(s, char)
		endPos += pos
		matchCount--
	}
	return endPos
}
