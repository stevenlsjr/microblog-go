package models

type Message struct {
	UuidModel
	AuthorId string
	Content  string
}
