package whois

import (
	"fmt"
	"strings"
)

type domain struct {
	name string
}

func (d domain) sld() string {
	s, err := d.splitName()
	if err != nil {
		return ""
	}
	return s[0]
}

func (d domain) tld() string {
	s, err := d.splitName()
	if err != nil {
		return ""
	}
	return s[1]
}

func (d domain) splitName() ([]string, error) {
	s := strings.SplitN(d.name, ".", 2)
	if len(s) != 2 {
		return nil, fmt.Errorf("%s is invalid name.", d.name)
	}
	return s, nil
}
