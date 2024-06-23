package models

type BaseResponseModels struct {
	CodeResponse  int         `json:"Code"`
	HeaderMessage string      `json:"HeaderMessage"`
	Message       string      `json:"message"`
	Data          interface{} `json:"Data"`
}