package model

import (
	"encoding/json"
	"strings"
)

type TypeMessage int

const (
	TYPEMESSAGE_CREATE_ERROR TypeMessage = 0
	TYPEMESSAGE_CREATE_ACCOUNT TypeMessage = 1
	TYPEMESSAGE_LOGIN TypeMessage = 2
	TYPEMESSAGE_CHAT TypeMessage = 3
)

type TypeData struct {
	TypeMsg TypeMessage
	Data    json.RawMessage
}

func (d *TypeData) GetType() TypeMessage {

	return d.TypeMsg
}

func (d *TypeData) SetType(typeMsg TypeMessage) {
	d.TypeMsg = typeMsg
}

func (d *TypeData) GetData() json.RawMessage {
	return d.Data
}

func (d *TypeData) GetValue(b []byte) (interface{}, error) {

	err := json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}
	//log.Println("type of incomming data:", d.GetType())
	//log.Println("fomat of incomming data", string(d.GetData()))

	return d.SelectType(d.GetType())
}
func (d *TypeData)SelectType(t TypeMessage) (interface{}, error) {
	switch t {
	case TYPEMESSAGE_CREATE_ACCOUNT:
		var p User
		err := json.Unmarshal([]byte(d.GetData()), &p)
		if err != nil {
			return nil, err
		}
		return &p, nil
	case TYPEMESSAGE_LOGIN:

	case TYPEMESSAGE_CHAT:

	}
	return nil, nil
}
