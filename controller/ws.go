package controller

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"log"
	"flag"
	"html/template"
	model "../model"
	"encoding/json"
)

var upgrader websocket.Upgrader
var (
	Addr = flag.String("addr", "127.0.0.1:9090", "http service address")
	CmdPath string
	homeTempl = template.Must(template.ParseFiles("../chat/view/Home.html"))
)

func wsInternalErrorPrint(ws websocket.Conn, msg string, err error) {
	//ws.WriteMessage(websocket.TextMessage, []byte("Internal server we error"))
	log.Println("ws:" + msg, err)
}
func checkOrigin(r *http.Request) bool {
	return true
}
func processData(obj interface{}, conn websocket.Conn) {

	var err error
	var notification string
	var jsType model.JsonType

	switch obj.(type) {
	case *model.User:

		user := obj.(*model.User)
		user.Println()
		err = user.Add()
		if err != nil {
			notification = err.Error()
			jsType.SetNotification(notification)
		} else {
			jsType.SetNotification("")
		}
		jsType.SetType(model.TYPEMESSAGE_CREATE_ACCOUNT)
	}
	if err != nil {
		wsInternalErrorPrint(conn, "read", err)
		return
	}
	ReplyClient(jsType, conn)
}
func ReplyClient(jsType model.JsonType, conn websocket.Conn) {

	// reply to client

	data, err := json.Marshal(&jsType)
	if err != nil {
		return
	}

	jsData := string(data)
	fmt.Println(jsData)

	err = conn.WriteMessage(websocket.TextMessage, []byte(jsData))
	if err != nil {
		wsInternalErrorPrint(conn, "write", err)
		return
	}
}

func ListenerIncomming(conn *websocket.Conn) {

	defer conn.Close()
	for {

		if r := recover(); r != nil {
			fmt.Println("Recoverd readMessage", r)
		}

		_, data, err := conn.ReadMessage()
		if err == nil {

			var typeData = model.TypeData{}
			obj, err := typeData.GetValue(data)

			if err == nil {
				processData(obj, *conn)
			}
		}
	}
}
func ServerWs(w http.ResponseWriter, r *http.Request) {

	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize:1024,
		CheckOrigin:checkOrigin,
	};

	var err error
	var conn *websocket.Conn
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	ListenerIncomming(conn)
}
