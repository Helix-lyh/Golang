package main

import (
	"bytes"
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(intToStr(values))
}

func intToStr(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
