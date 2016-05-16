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
func processData(obj interface{}, conn websocket.Conn) {

	var err error
	var notification string

	switch obj.(type) {
	case *model.User:

		user := obj.(*model.User)
		user.Println()
		err = user.Add()
		if err != nil {
			notification = err.Error()
		}
	}
	if err != nil {
		wsInternalErrorPrint(conn, "read", err)
		return
	}
	ReplyClient(notification, conn)
}
func ReplyClient(notification string, conn websocket.Conn) {

	// reply to client

	errType := &model.ErrorType{}

	if len(notification) > 0 {
		errType.SetError(true)
		errType.SetNotification(notification)
	} else {
		errType.SetError(false)
		errType.SetNotification("")
	}

	jsData, err := json.Marshal(errType)
	if err != nil {
		return
	}

	err = conn.WriteMessage(NoFrame, []byte( string(jsData)))
	if err != nil {
		wsInternalErrorPrint(conn, "write", err)
		return
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
				processData(obj, &conn)
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
	var conn *websocket.Conn
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	ListenerIncomming(conn)
}
