package main

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String()string  {
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request){
	for item, price := range db{
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item:%q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// ListenAndServe 中第二个参数的接口 要求接口器提供ServerHTTP 方法
func main()  {
	db := database{"shoes":50, "socks": 5}
	fmt.Println(fmt.Sprintf("listen on %s:\n", "8000"))
	// 这种是利用的
	http.HandleFunc("/list", db.list)
	http.ListenAndServe("localhost:8000", nil)
}