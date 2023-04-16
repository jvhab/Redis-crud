package model

import (
	"github.com/google/uuid"
	"time"
)

type News struct {
	Id        uuid.UUID `json:"id" db:"id"`
	AuthorId  uuid.UUID `json:"author_id" db:"author_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Category  string    `json:"category" db:"category"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
