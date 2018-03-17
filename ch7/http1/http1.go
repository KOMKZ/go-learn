package main

import (
	"fmt"
	"net/http"
	"log"
)

type dollars float32

func (d dollars) String()string  {
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollars


func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request){
	for item,price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// ListenAndServe 中第二个参数的接口 要求接口器提供ServerHTTP 方法
func main()  {
	db := database{"shoes":50, "socks": 5}
	fmt.Println(fmt.Sprintf("listen on %s:\n", "8000"))
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}