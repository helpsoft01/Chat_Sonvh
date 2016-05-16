package controller

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"log"
	"flag"
	"html/template"
	model "../model"
)

var upgrader websocket.Upgrader
var (
	Addr = flag.String("addr", "127.0.0.1:9090", "http service address")
	CmdPath string
	homeTempl = template.Must(template.ParseFiles("../chat/view/Home.html"))
	conn *websocket.Conn
)

const (
	NoFrame = -1
)

func wsInternalErrorPrint(ws *websocket.Conn, msg string, err error) {
	ws.WriteMessage(websocket.TextMessage, []byte("Internal server we error"))
	log.Println("ws:" + msg, err)
}
func checkSameOrigin(r *http.Request) bool {
	return true
}
func processData(obj interface{}) {

	var err error
	var notification string
	switch obj.(type) {
	case *model.User:

		user := obj.(*model.User)
		user.Println()
		err = user.Add()
		if err != nil {
			notification = err
		}
	}

	if err != nil {
		wsInternalErrorPrint(conn, "read", err)
	}

	// reply to client
	err = conn.WriteMessage(NoFrame, notification)
	if err != nil {
		wsInternalErrorPrint(conn, "write", err)
	}
}
func ListenerIncomming(conn *websocket.Conn) {

	defer conn.Close()
	for {
		_, data, err := conn.ReadMessage()
		if err == nil {

			var typeData = model.TypeData{}
			obj, err := typeData.GetValue(data)

			if err == nil {
				processData(obj)
			}
		}
	}
}
func ServerWs(w http.ResponseWriter, r *http.Request) {

	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize:1024,
		CheckOrigin:checkSameOrigin,
	};

	var err error
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	go ListenerIncomming(conn)
}
