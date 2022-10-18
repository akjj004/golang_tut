package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var items = [3]item{
	{"OnePlus 10 Pro", 4500},
	{"Asus X513EA", 2000},
	{"Samsung UE55TU7022K", 1900},
}

func getItemID(url string) int {
	fields := strings.Split(url, "/")
	if len(fields) < 2 || fields[2] == "" {
		return -1
	} // - /item/
	id, err := strconv.Atoi(fields[2])
	if err != nil {
		return -2
	} // - /item/not_a_number
	return id

}

// W pliku main.go utworzyć strukturę item
// zawierającą następujące pola:
// Name (string), Price (float)

type item struct {
	Name  string
	Price float32
}

// Utworzyć serwer http obsługujący
// stronę /item/ przez funkcję itemFunc.

func itemFunc(w http.ResponseWriter, r *http.Request) {
	// pobranie ID przedmiotu z adresu strony
	id := getItemID(r.RequestURI)
	fmt.Println(id)
	if id < 0 {
		// zwrócenie statycznej strony strona.html
		http.ServeFile(w, r, "TASKC4T2/pages/strona.html")
	}

	if id >= 0 {
		tmpl, _ := template.ParseFiles("TASKC4T2/pages/parse.html")
		tmpl.Execute(w, &items[id])
		file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0644)
		file.WriteString(fmt.Sprint(id) + "\t" + time.Now().String() + "\n")
		file.Close()

	}

}

func main() {
	http.HandleFunc("/item/", itemFunc)
	// http.HandleFunc("/parse/", itemFunc)
	http.ListenAndServe("localhost:8080", nil)
}
