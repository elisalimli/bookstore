package models

import "github.com/google/uuid"

type Post struct {
	// gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Title string
	Body  string
}
