package models

type Ping struct {
	Id      int    `json:"id"`
	Message string `json:"msg"`
	Gender  string `json:"gender"`
}
