package entity

import (
	"time"
)

type Todo struct {
	Id        uint64 		`json:"id"`
	Title     string 		`json:"title"`
	IsDone    bool 			`json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}