package book

import "rest-api-learning/internal/author"

type Book struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Authors []author.Author `json:"authors"`
}
