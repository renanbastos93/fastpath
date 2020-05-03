package fastpath

import (
	"strings"
)

type seg struct {
	Param      string
	Const      string
	ConstLen   int
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
	patternCount := len(aPattern)
	var out = make([]seg, patternCount)
	var params []string
	for i := 0; i < patternCount; i++ {
		partLen := len(aPattern[i])
		if partLen == 0 { // skip empty parts
			continue
		}
		// is parameter
		if aPattern[i][0] == '*' || aPattern[i][0] == ':' {
			out[i] = seg{
				Param:      paramTrimmer(aPattern[i]),
				IsParam:    true,
				IsOptional: aPattern[i] == "*" || aPattern[i][partLen-1] == '?',
			}
			params = append(params, out[i].Param)
		} else {
			out[i] = seg{
				Const:    aPattern[i],
				ConstLen: len(aPattern[i]),
			}
		}
	}
	// TODO: optimize it
	lastIndex := -1
	reducedOut := make([]seg, patternCount)
	for _, seg := range out {
		if lastIndex != -1 && reducedOut[lastIndex].ConstLen > 0 && seg.ConstLen > 0 {
			reducedOut[lastIndex].Const += "/" + seg.Const
			reducedOut[lastIndex].ConstLen += 1 + seg.ConstLen
			continue
		} else {
			lastIndex++
		}
		reducedOut[lastIndex] = seg
	}
	if lastIndex != -1 {
		reducedOut[lastIndex].IsLast = true
	}
	p = Path{Segs: reducedOut[:lastIndex+1], Params: params}
	return
}

// Match ...
func (p *Path) Match(s string) ([]string, bool) {
	params := make([]string, len(p.Params), cap(p.Params))
	var i, j, paramsIterator, partLen int
	if len(s) > 0 {
		s = s[1:]
	}
	for index, segment := range p.Segs {
		partLen = len(s)
		if segment.IsParam {
			if segment.IsLast {
				i = partLen
			} else if !segment.IsLast && segment.Param == "*" {
				// for the expressjs behavior -> "/api/*/:param" - "/api/joker/batman/robin/1" -> "joker/batman/robin", "1"
				i = findCharPos(s, '/', strings.Count(s, "/")-(len(p.Segs)-(index+1))+1)
			} else {
				i = strings.IndexByte(s, '/')
			}
			if i == -1 {
				i = partLen
			}
			if !segment.IsOptional && s == "" {
				return nil, false
			}
			params[paramsIterator] = s[:i]
			paramsIterator++
		} else {
			i = segment.ConstLen
			if partLen < i || (i == 0 && partLen > 0) || s[:i] != segment.Const {
				return nil, false
			}
		}

		if partLen > 0 {
			j = i + 1
			if segment.IsLast || partLen < j {
				j = i
			}

			s = s[j:]
		}
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
