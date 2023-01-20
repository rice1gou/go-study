package main

import (
	"go-study/pkg/lg"
	"net/http"
)

func main() {
	l := lg.LoggerAdapter(lg.LogOutput)
	ds := lg.NewSimpleDataStore()
	logic := lg.NewSimpleLogic(l, ds)
	c := lg.NewController(l, logic)
	http.HandleFunc("/hello", c.HandleGreeting)
	http.ListenAndServe(":8080", nil)
}
