package whois

import "strings"

type Response interface {
	Raw() string
	Exist() bool
}

type base struct {
	res string
}

func (b base) Raw() string {
	return b.res
}

type verisign struct {
	base
}

func (v verisign) Exist() bool {
	return !strings.Contains(v.res, "No match for ")
}

type jprs struct {
	base
}

func (j jprs) Exist() bool {
	return !strings.Contains(j.res, "No match!!")
}

type gmo struct {
	base
}

func (g gmo) Exist() bool {
	return !strings.Contains(g.res, "DOMAIN NOT FOUND")
}

func newResponse(tld, response string) Response {
	b := base{response}
	switch tld {
	case "com", "net":
		return verisign{base: b}
	case "jp":
		return jprs{base: b}
	case "tokyo":
		return gmo{base: b}
	}
	return nil
}
