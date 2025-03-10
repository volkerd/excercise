package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func serve() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/news", handleAll)
	r.HandleFunc("/news/{id}", handleSingle)
	r.HandleFunc("/api/news", handleAllApi).Methods("GET")
	r.HandleFunc("/api/news/{id}", handleSingleApi).Methods("GET")
	go log.Fatal(http.ListenAndServe(":8080", r))
}

func handleAll(w http.ResponseWriter, r *http.Request) {
	log.Println("handleAll called")
	fmt.Fprint(w, "<h1>News von heise.de</h1>")
	news := newsList(5)
	for i := 0; i < len(news.News); i++ {
		item := news.News[i]
		fmt.Fprintf(w, `
		<hr/>
		<h2><a href='news/`+item.Id+`'>%v</a></h2>
		<h3>Veröffentlicht am %v</h3>
		<p>%v</p>
		`, item.Title, item.Id, formatDate(item.Published), item.Description)
	}
}

func handleSingle(w http.ResponseWriter, r *http.Request) {
	log.Println("handleSingle called")
	vars := mux.Vars(r)
	id := vars["id"]
	news, err := newsSingle(id)
	if err != nil {
		fmt.Fprintln(w, "not found")
		return
	}
	fmt.Fprintf(w, `
		<h1>%v (%v)</h1>
		<h3>Veröffentlicht am %v</h3>
		<p>%v %v</p>
		`, news.Title, id, formatDate(news.Published), news.Description, news.Id)
}

func handleAllApi(w http.ResponseWriter, r *http.Request) {
	log.Println("handleAllApi called")
	js, _ := json.Marshal(newsList(5))
	w.Write(js)
}

func handleSingleApi(w http.ResponseWriter, r *http.Request) {
	log.Println("handleSingleApi called")
	vars := mux.Vars(r)
	id := vars["id"]
	news, err := newsSingle(id)
	if err != nil {
		fmt.Fprintln(w, "not found")
		return
	}
	js, _ := json.Marshal(news)
	w.Write(js)
}
