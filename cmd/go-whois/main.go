package main

import (
	"fmt"
	"os"

	"github.com/monochromegane/go-whois"
)

func main() {
	fmt.Printf("Hello, world\n")
	whois.Query(os.Args[1])
}
