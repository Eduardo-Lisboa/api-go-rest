package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"name"`
	Historia string `json:"history"`
}

var Personalidades []Personalidade
