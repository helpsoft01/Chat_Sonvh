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
	"io"
	"io/ioutil"
)

var upgrader websocket.Upgrader
var (
	Addr = flag.String("addr", "127.0.0.1:9090", "http service address")
	CmdPath string
	homeTempl = template.Must(template.ParseFiles("../chat/view/Home.html"))
)

func wsInternalErrorPrint(msg string, err error) {
	//ws.WriteMessage(websocket.TextMessage, []byte("Internal server we error"))
	log.Println("ws:" + msg, err)
}
func checkOrigin(r *http.Request) bool {
	return true
}
func processData(typeMsg model.TypeMessage, obj interface{}, conn websocket.Conn) {

	var err error
	var notification string
	var jsType model.JsonType

	switch obj.(type) {
	case *model.User:

		switch typeMsg {
		case model.TYPEMESSAGE_CREATE_ACCOUNT:

			user := obj.(*model.User)
			user.Println()
			err = user.Add()

			if err != nil {
				notification = err.Error()
				jsType.SetNotification(notification)
			} else {
				jsType.SetNotification("")
			}

		case model.TYPEMESSAGE_LOGIN:

			user := obj.(*model.User)
			user.Println()
			if user.CheckExistByName() {

				jsType.SetNotification("")
			} else {

				jsType.SetNotification("Not Exist User")
			}
		}
	}

	jsType.SetType(typeMsg)
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
		wsInternalErrorPrint("write", err)
		return
	}
}
func ReadMessage(c *websocket.Conn) (messageType int, p []byte, err error) {

	var r io.Reader
	messageType, r, err = c.NextReader()
	if err != nil {
		return messageType, nil, err
	}
	p, err = ioutil.ReadAll(r)
	return messageType, p, err
}
func ListenerIncomming(conn *websocket.Conn) {

	defer conn.Close()
	for {
		if r := recover(); r != nil {
			fmt.Println("Recoverd readMessage", r)
		}

		_, data, err := ReadMessage(conn)

		if err == nil {
			var objData = model.TypeData{}
			obj, err := objData.GetValue(data)

			if err == nil {
				processData(objData.GetType(), obj, *conn)
			}
		} else {
			fmt.Println("ws close:", err)
			if c, k := err.(*websocket.CloseError); k {
				if c.Code == websocket.CloseGoingAway {
					//fmt.Println("ws close:", err)
				}
			}
		}

		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "woops"))

	}
}
func WsHandler(w http.ResponseWriter, r *http.Request) {

	upgrader = websocket.Upgrader{
		ReadBufferSize: 2 * 1024,
		WriteBufferSize:2 * 1024,
		CheckOrigin:checkOrigin,
	};

	var err error
	var conn *websocket.Conn
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	go ListenerIncomming(conn)
}
