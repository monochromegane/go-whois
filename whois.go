package whois

import "net"

func Query(d string) (Response, error) {
	domain := domain{name: d}
	server, err := findServer(domain.tld())
	if err != nil {
		return nil, err
	}
	res, err := query(server, domain)
	return res, err
}

func query(server string, domain domain) (Response, error) {

	conn, err := net.Dial("tcp", server+":43")
	if err != nil {
		return nil, err
	}
	conn.Write([]byte(domain.name + "\r\n"))
	buf := make([]byte, 1024)
	res := []byte{}
	for {
		numbytes, err := conn.Read(buf)
		sbuf := buf[0:numbytes]
		res = append(res, sbuf...)
		if err != nil {
			break
		}
	}
	conn.Close()
	return newResponse(domain.tld(), string(res)), nil
}
