package main

import (
	"fmt"
	"os"

	"github.com/monochromegane/go-whois"
)

var usage = `usage:
  go-whois DOMAIN_NAME`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	res, err := whois.Query(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Raw())
}
