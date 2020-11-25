package main

import (
	"fmt"
	"log"
	"net/http"
)
func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "/Users/shubham.sethi/work/personal/go-programs/http/index.html"
	}
	fmt.Println(p)
	http.ServeFile(w, r, p)
}
func main() {

	PORT:= ":8085"

	http.HandleFunc("/", serveFiles)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println(err)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

}
