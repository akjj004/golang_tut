package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	// zwrócenie strony głównej
	fmt.Fprintf(w, "<html><body>STRONA GŁÓWNA</body></html>")
}
func itemFunc(w http.ResponseWriter, r *http.Request) {
	// zwrócenie strony /item/* (* - oznacza dowolny ciąg znaków)
	fmt.Fprintf(w, "<html><body>STRONA ITEM<br>ADRES: ")
	fmt.Fprintf(w, r.RequestURI)
	fmt.Fprintf(w, "<br>METODA: ")
	fmt.Fprintf(w, r.Method)
	fmt.Fprintf(w, "</body></html>")
}

func stronaFunc(w http.ResponseWriter, r *http.Request) { // zwrócenie statycznej strony strona.html http.ServeFile(w, r, "pages/strona.html")
	http.ServeFile(w, r, "TASKC4/pages/strona.html")
}

type student struct {
	Name  string
	Index int
}

func parseFunc(w http.ResponseWriter, r *http.Request) { // zwrócenie strony o dynamicznej zawartości
	tmpl, _ := template.ParseFiles("TASKC4/pages/parse.html")
	tmpl.Execute(w, &student{"Jan", 12345})
}

func main() {
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/item/", itemFunc)
	http.HandleFunc("/strona/", stronaFunc)
	http.HandleFunc("/parse/", parseFunc)
	http.ListenAndServe("localhost:8080", nil)
}
