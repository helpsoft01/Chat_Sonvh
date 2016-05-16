package model

type ErrorType struct {
	IsError      bool
	Notification string
}

func (e *ErrorType) GetError() bool {
	return e.IsError
}
func (e *ErrorType) SetError(isErr bool) {
	e.IsError = isErr
}
func (e *ErrorType) GetNotification() string {
	return e.Notification
}
func (e *ErrorType) SetNotification(notification string) {
	e.Notification = notification
}
