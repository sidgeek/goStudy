package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhelloName)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!") //输出到客户端的信息
}
