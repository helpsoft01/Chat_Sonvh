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
	http.HandleFunc("/ws", ws.ServerWs)
	//http.HandleFunc("/", ws.ServerHome)
	log.Fatal(http.ListenAndServe(*ws.Addr, nil))

	//var input string
	//fmt.Scan(&input)
}

