package models

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
