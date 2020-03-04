package main

import (
	"fmt"
	"strings"
)

func main() {
	helloworld := "hello, world"
	HelloCamel := strings.Title(helloworld)

	HelloCapital := strings.ToUpper(helloworld)

	fmt.Println(HelloCamel)
	fmt.Println(HelloCapital)
}
