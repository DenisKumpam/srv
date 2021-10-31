package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "hello\n")
}

func readFile(w http.ResponseWriter, req *http.Request)  {
	text := os.ReadFile("file.txt")
	fmt.Fprintf(w, text,"\n")
}

func headers(w http.ResponseWriter, req *http.Request)  {
	for name, headers := range req.Header {
		for _, h := range headers{
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
