package main

import "fmt"

func main() {
	fmt.Println(comma("hello world"))
}
func comma(s string) string {
	fmt.Println(s)
	n := len(s)
	if n < 2 {
		return s
	}
	// 递归的用法
	return comma(s[:n-1]) + "," + s[n-1:]
}
