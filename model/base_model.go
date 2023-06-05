package model

import (
	"time"
)

type BaseModel struct {
	ID        string    `gorm:"primaryKey type:uuid;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
