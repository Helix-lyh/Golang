package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "-"
	}
	fmt.Println("---方案1 ---")
	fmt.Println(s)
	fmt.Println("---方案2 ---")
	fmt.Println(strings.Join(os.Args[1:], "-"))
	fmt.Println("---方案3 ---")
	fmt.Println(os.Args[1:])
}
