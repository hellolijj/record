package main

import (
	"net/http"
	"fmt"
)


type database map[string]float32

var db = database{"shoes": 32, "socks": 102, "lijj": 2321}


func (db database)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.URL.Path {

	case "/list":
		for item, price := range db{
			fmt.Fprintf(w, "%s: %f\n", string(item), float32(price))
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %v", item)
			return
		}
		fmt.Fprintln(w, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "no such pages")
	}
}

func main() {
	http.ListenAndServe("localhost:8000", db)
}