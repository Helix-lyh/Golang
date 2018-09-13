package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args {
		fmt.Printf("fileName:fileName=%s\n", fileName)
		data, err := ioutil.ReadFile("E:\\golangStudy\\src\\awesomeProject\\fordup")
		num := 0
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v \n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			num++
		}
		for line, n := range counts {
			if n > 0 {
				// 此处踩了个坑 因为文件中包含 \r 换行符 导致输出时异常，需要手动将 line 中的 \r 进行去除
				fmt.Printf("字符串 %s 出现了 %d次\n", strings.TrimSuffix(line, "\r"), n)
			}
		}
		fmt.Printf("字符串总行数%d", num)
	}

}
