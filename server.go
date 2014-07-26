package whois

import "fmt"

var servers = map[string]string{
	"com":   "whois.verisign-grs.com",
	"net":   "whois.verisign-grs.com",
	"jp":    "whois.jprs.jp",
	"tokyo": "whois.nic.tokyo",
}

func findServer(tld string) (string, error) {
	server, exist := servers[tld]
	if !exist {
		return "", fmt.Errorf("%s's whois server dosen't register.", tld)
	}
	return server, nil
}
