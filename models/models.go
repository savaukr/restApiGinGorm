package model

type Users struct {
	Id uint
	UserName string
}
type Messages struct {
	Id uint
	NameSender string
	NameReciever string
	Message string
}