package main

import (
	ws "../controller"
	"flag"
	"log"
	"net/http"
)

func main() {

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", ws.WsHandler)
	//http.HandleFunc("/", ws.ServerHome)
	log.Fatal(http.ListenAndServe(*ws.Addr, nil))

	//var input string
	//fmt.Scan(&input)
}

