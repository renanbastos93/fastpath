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
	var patternCount int
	aPattern := []string{""}
	if pattern != "" {
		aPattern = strings.Split(pattern, "/")[1:] // every route starts with an "/"
	}
	patternCount = len(aPattern)

	var out = make([]seg, patternCount)
	var params []string
	var segIndex int
	for i := 0; i < patternCount; i++ {
		partLen := len(aPattern[i])
		if partLen == 0 { // skip empty parts
			continue
		}
		// is parameter ?
		if aPattern[i][0] == '*' || aPattern[i][0] == ':' {
			out[segIndex] = seg{
				Param:      paramTrimmer(aPattern[i]),
				IsParam:    true,
				IsOptional: aPattern[i] == "*" || aPattern[i][partLen-1] == '?',
			}
			params = append(params, out[segIndex].Param)
		} else {
			// combine const segments
			if segIndex > 0 && out[segIndex-1].IsParam == false {
				segIndex--
				out[segIndex].Const += "/" + aPattern[i]
				// create new const segment
			} else {
				out[segIndex] = seg{
					Const: aPattern[i],
				}
			}
		}
		segIndex++
	}
	if segIndex == 0 {
		segIndex++
	}
	out[segIndex-1].IsLast = true

	p = Path{Segs: out[:segIndex:segIndex], Params: params}
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
		// check parameter
		if segment.IsParam {
			// determine parameter length
			if segment.IsLast {
				i = partLen
			} else if segment.Param == "*" {
				// for the expressjs behavior -> "/api/*/:param" - "/api/joker/batman/robin/1" -> "joker/batman/robin", "1"
				i = findCharPos(s, '/', strings.Count(s, "/")-(len(p.Segs)-(index+1))+1)
			} else {
				i = strings.IndexByte(s, '/')
			}
			if i == -1 {
				i = partLen
			}

			if false == segment.IsOptional && i == 0 {
				return nil, false
			}

			params[paramsIterator] = s[:i]
			paramsIterator++
		} else {
			// check const segment
			i = len(segment.Const)
			if partLen < i || (i == 0 && partLen > 0) || s[:i] != segment.Const {
				return nil, false
			}
		}

		// reduce founded part from the string
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
