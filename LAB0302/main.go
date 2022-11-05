package main

import (
	"fmt"
	"lab0302/knn"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var recsys *knn.Knn

type  Data struct {
	Beer knn.Beer
	StyleName string
}

func beerFunc(w http.ResponseWriter, r *http.Request) {
	tenbeers := recsys.Get10RandomBeers()
	tmpl, _ := template.ParseFiles("pages/beer.html")
	tmpl.Execute(w, tenbeers)
}

func rekoFunc(w http.ResponseWriter, r *http.Request) {
 // gdy metoda inna niż POST - error 403
 if r.Method != "POST" {
 http.Error(w, "Only POST supported", http.StatusForbidden)
 return
 }
 // parsowanie danych z formularza
 r.ParseForm()
 // iteracja danych z formularza i jednoczesna ocena piw
 for name, element := range r.PostForm {
 rate, _ := strconv.ParseFloat(element[0], 64) // parsowanie
 id, _ := strconv.Atoi(name[4:]) // wycięcie id
 beer := recsys.GetBeerByID(id) // pobranie piwa
 beer.Rate = rate // ocena piwa
 fmt.Println("Name:", name, ", Id:", id, ", Rate:", rate)
 }
//  // pobranie rekomendacji
 recbeers := recsys.GetRecommendation()
 tmpl, _ := template.ParseFiles("pages/reko.html")
 tmpl.Execute(w, recbeers)
}



func getItemID(url string) int {
fields := strings.Split(url, "/")
if len(fields) < 2 || fields[2] == "" { return -1 } // - /item/
id, err := strconv.Atoi(fields[2])
if err != nil { return -2 } // - /item/not_a_number
return id // /item/nr/
}

func beerooFunc(w http.ResponseWriter, r *http.Request){
	id := getItemID(r.RequestURI)
	beroo := recsys.GetBeerByID(id) 
	//  style := recsys.GetStyleName(beroo.Style)
	style := recsys.GetStyleName(beroo.Style)
	tmpl, _ := template.ParseFiles("pages/beroo.html")
	tmpl.Execute(w, beroo)
	tmpl.Execute(w, &Data{*beroo , style})

}

func main() {
	// po testach działania należy te linijki zostawić
	rand.Seed(time.Now().UnixMilli())
	recsys = knn.Initialize()
	http.HandleFunc("/beer/", beerFunc)
	http.HandleFunc("/reko/", rekoFunc)
	http.HandleFunc("/beroo/", beerooFunc)
	// http.HandleFunc("/parse/", itemFunc)
	http.ListenAndServe("localhost:8081", nil)

	// a te poniżej zakomentować
	// rand.Seed(time.Now().UnixMilli())
	// recsys = knn.Initialize()
	// var rateText string
	// var rateCount int = 0
	// for rateCount < 10 {
	// 	beer := recsys.GetRandomBeer()
	// 	beer.DisplayInformation(recsys)
	// 	fmt.Println("Ocen piwo (1-5): ")
	// 	fmt.Scanln(&rateText)
	// 	rate, error := strconv.ParseFloat(rateText, 64)
	// 	if error == nil && rate >= 1 && rate <= 5 {
	// 		beer.Rate = rate
	// 		rateCount++
	// 	}
	// }
	// reco := recsys.GetRecommendation()
	// fmt.Println("Rekomendowane piwa: ")
	// reco[0].DisplayInformation(recsys)
	// reco[1].DisplayInformation(recsys)
	// reco[2].DisplayInformation(recsys)
}
