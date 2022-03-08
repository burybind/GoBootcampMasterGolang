package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)

	excl := strings.Repeat("!", l)
	s := excl + msg + excl
	s = strings.ToUpper(s)

	fmt.Println(s)
}

