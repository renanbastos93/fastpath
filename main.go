package main

import (
	"fmt"
	"strings"
)

type seg struct {
	Param      string
	Const      string
	IsParam    bool
	IsOptional bool
	IsWildcard bool
}

type match struct {
	S      []seg
	Params map[string]string
}

func makePattern(pattern string) (m match) {
	if pattern == "*" {
		return
	}
	aPattern := strings.Split(pattern, "/")
	var i int = 0
	var out = make([]seg, len(aPattern))
	for i = 0; i < len(aPattern); i++ {
		if strings.HasPrefix(aPattern[i], ":") {
			if strings.HasSuffix(aPattern[i], "?") {
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
			out[i] = seg{
				Param:      aPattern[i],
				IsWildcard: true,
			}
		} else {
			out[i] = seg{
				Const: aPattern[i],
			}
		}
	}
	m = match{S: out[1:]}
	return
}

func (m *match) matchs(uri string) (res map[string]string, ok bool) {
	aURI := strings.Split(uri, "/")
	if len(aURI[1:]) > len(m.S) {
		return
	}
	res = map[string]string{}
	for k, v := range m.S {
		val := aURI[k+1]
		if v.IsParam && !v.IsOptional {
			if val == "" {
				fmt.Println("ERR: not match")
				return
			}
			res[v.Param[1:]] = val
		} else if v.IsParam && v.IsOptional {
			res[v.Param[1:len(v.Param)-1]] = val
		} else if v.IsWildcard {
			res[v.Param] = val
		}
	}
	ok = true
	return
}

func main() {
	str := "/api/:param/:opt?"
	x := makePattern(str)
	fmt.Println(x.matchs("/api/a/b"))
}
