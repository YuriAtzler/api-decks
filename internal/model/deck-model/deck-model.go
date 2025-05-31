package deckmodel

type ListDeck struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}


type CreateDeck struct {
	Name string `json:"name" binding:"required"`
}