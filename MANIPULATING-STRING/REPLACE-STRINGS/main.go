package main

import (
	"fmt"
	"strings"
)

func main() {
	helloworld := "hello, world"
	HelloEliarms := strings.Replace(helloworld, "world", "Eliarms", 1)
	fmt.Printf(HelloEliarms)
}
