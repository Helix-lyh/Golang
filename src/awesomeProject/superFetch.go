package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Printf("url = %v", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "superFetch: %v\n", err)
			os.Exit(1)
		}
		text, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "superFetch:reding %s, %v", url, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%d \n", text)
		fmt.Fprintf(os.Stdout, "%s", resp.Status)
	}
}
