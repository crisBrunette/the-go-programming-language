package main

import (
	"fmt"
	"net/http"

	"github.com/crisBrunette/the-go-programming-language/hello-world/responder"
	"rsc.io/quote"
)

func endPoint1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Well done!")
}
func endPoint2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, responder.HeaderWriter(req))
}

func main() {

	fmt.Println(quote.Hello())
	http.HandleFunc("/endPoint1", endPoint1)
	http.HandleFunc("/endPoint2", endPoint2)

	http.ListenAndServe(":8090", nil)
}
