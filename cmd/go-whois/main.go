package main

import (
	"fmt"
	"os"

	"github.com/monochromegane/go-whois"
)

func main() {
	res, err := whois.Query(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Raw())
}
