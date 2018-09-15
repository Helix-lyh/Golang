package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("进入handler2 count = %d", count)
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "url path = %q\n", r.URL.Path)
	fmt.Printf("退出handler2 count = %d", count)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count = %d", count)
	mu.Unlock()
}
