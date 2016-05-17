package model

type JsonType struct {
	TypeMsg      TypeMessage
	Notification string
}

func (d *JsonType) GetType() TypeMessage {

	return d.TypeMsg
}

func (d *JsonType) SetType(typeMsg TypeMessage) {
	d.TypeMsg = typeMsg
}

func (e *JsonType) GetNotification() string {
	return e.Notification
}
func (e *JsonType) SetNotification(notification string) {
	e.Notification = notification
}
