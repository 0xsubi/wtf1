package models

type RadioMessage struct {
	Uid string `json:"uid"`
	Message string `json:"message"`
	Color string `json:"color"`
	Speaker string `json:"speaker"`
	Event string `json:"event"`
	Link string `json:"link"`
}
