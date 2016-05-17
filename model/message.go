package model

type Message struct {
	IdSender   int64
	IdReceiver int64
	Text       string
}

func (m *Message) GetIdSender() int64 {
	return m.IdSender
}
func (m *Message) SetIdSender(id int64) {
	m.IdSender = id
}
func (m *Message) GetIdReceiver() int64 {
	return m.IdReceiver
}
func (m *Message) SetIdReceiver(id int64) {
	m.IdReceiver = id
}
func (m *Message) GetText() string {
	return m.Text
}
func (m *Message) SetText(txt string) {
	m.Text = txt
}