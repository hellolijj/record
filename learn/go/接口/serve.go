package main

import (
	"net/http"
	"fmt"
	"strconv"
)


type database map[string]float32

var db = database{"shoes": 32, "socks": 102}

func (db database)list(w http.ResponseWriter, r *http.Request)  {
	for item, price := range db{
		fmt.Fprintf(w, "%s: %f\n", string(item), float32(price))
	}
}

func (db database)price(w http.ResponseWriter, r *http.Request)  {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %v", item)
		return
	}
	fmt.Fprintln(w, price)
}


func (db database)add(w http.ResponseWriter, r *http.Request)  {
	//http://127.0.0.1:8001/add?item=add&price=23

	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")


	price32, _ := strconv.ParseFloat(price, 32)
	db[item] = float32(price32)

	for item, price := range db{
		fmt.Fprintf(w, "%s: %f\n", string(item), float32(price))
	}
}



func (db database)delete(w http.ResponseWriter, r *http.Request)  {
	//http://127.0.0.1:8001/delete?item=add&price=23

	item := r.URL.Query().Get("item")
	delete(db, item)
	for item, price := range db{
		fmt.Fprintf(w, "%s: %f\n", string(item), float32(price))
	}
}

func (db database)update(w http.ResponseWriter, r *http.Request)  {
	//http://127.0.0.1:8001/update?item=add&price=23

	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")


	price32, _ := strconv.ParseFloat(price, 32)
	db[item] = float32(price32)

	for item, price := range db{
		fmt.Fprintf(w, "%s: %f\n", string(item), float32(price))
	}
}

//设计一个服务器实现增删改查
func main() {
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/delete", db.update)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.ListenAndServe("localhost:8001", nil)

}