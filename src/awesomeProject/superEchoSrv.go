package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe("0.0.0.0:8090", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method = %s URL = %s  Proto=%s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host = %q RemotAddr = %q\n", r.Host, r.RemoteAddr)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemotAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm异常")
	}
	fmt.Fprintf(w, "Form Data= %+v\n", r.Form)

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	bodyStr, err := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "bodyStr异常")
	}
	fmt.Fprintf(w, "Body = %q\n", bodyStr)
}
