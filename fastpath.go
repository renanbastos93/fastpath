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
	var out []seg
	var params []string
	startSignPos := 0
	var specialChar int32 = 0

	lastPos := len(pattern) - 1
	for pos, char := range pattern {
		if specialChar != 0 && '/' == char {
			out = append(out, seg{
				Param:      paramTrimmer(pattern[startSignPos:pos]),
				IsParam:    true,
				IsOptional: pattern[pos-1] == '?' || specialChar == '*',
			})
			params = append(params, paramTrimmer(pattern[startSignPos:pos]))
			startSignPos = pos
			specialChar = 0
		} else if '*' == char || ':' == char {
			out = append(out, seg{
				Const:       pattern[startSignPos:pos],
				ConstLength: len(pattern[startSignPos:pos]),
			})
			startSignPos = pos
			specialChar = char
		}

		if lastPos == pos {
			if specialChar != 0 {
				out = append(out, seg{
					Param:      paramTrimmer(pattern[startSignPos : pos+1]),
					IsParam:    true,
					IsOptional: pattern[pos] == '?' || specialChar == '*',
				})
				params = append(params, paramTrimmer(pattern[startSignPos:pos+1]))
			} else {
				out = append(out, seg{
					Const:       pattern[startSignPos : pos+1],
					ConstLength: len(pattern[startSignPos : pos+1]),
				})
			}
		}
	}
	p = Path{Segs: out, SegsLength: len(out), Params: params}
	return
}

// Match ...
func (p *Path) Match(s string) ([]string, bool) {
	params := make([]string, len(p.Params), cap(p.Params)) // reuse slice ?
	originalLen := len(s)
	paramsIterator := 0

	for index, segment := range p.Segs {

		if s == "" {
			if segment.IsOptional == true {
				continue
			} else if p.SegsLength-2 == index && p.Segs[index+1].Const == "" && p.Segs[index+1].IsOptional == true {
				break
			} else if p.SegsLength-1 == index && segment.Const == "/" { // not strict
				break
			}
			return nil, false
		}

		if segment.ConstLength != 0 { // is const part
			if len(s) < segment.ConstLength {
				return nil, false
			}
			if s[:segment.ConstLength] != segment.Const && s[:segment.ConstLength-1] != segment.Const {
				return nil, false
			}
			s = s[segment.ConstLength:]
		} else {
			if segment.Param == "*" {
				params[paramsIterator] = s
				paramsIterator++
				s = ""
			} else {
				nextPart := strings.IndexByte(s, '/')
				if nextPart == -1 {
					nextPart = len(s)
				}
				params[paramsIterator] = s[:nextPart]
				paramsIterator++
				s = s[nextPart:]
			}
		}
	}
	if s == "" && originalLen != 0 {
		return params, true
	}

	return nil, false
}

func paramTrimmer(param string) string {
	start := 0
	end := len(param) - 1

	if param[start] == 58 { // is :
		start++
	}
	if param[end] != 63 { // is ?
		end++
	}

	return param[start:end]
}
