package main

import (
	ws "../controller"
	"flag"
	"log"
	"net/http"
	flatbuffers "github.com/google/flatbuffers/go"
	helper "../helper"
	"fmt"
)

func main_() {

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", ws.WsHandler)
	//http.HandleFunc("/", ws.ServerHome)
	log.Fatal(http.ListenAndServe(*ws.Addr, nil))

	//var input string
	//fmt.Scan(&input)
}

func main(){
	b:=flatbuffers.NewBuilder(0)
	buf:= helper.MakeUser(b,[]byte("Ố trời"),42)
	name,id:= helper.ReadUser(buf)

	fmt.Printf("%s has %d . The encoded data is %d byte long",name,id,len(buf))
}

