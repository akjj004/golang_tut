package main

import (
	"fmt"
	"net/http"
)

func funcStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>OK HELLO!<body></html>")

}

func main() {

	http.HandleFunc("/start/", funcStart)
	http.ListenAndServe(":8080", nil)
}
