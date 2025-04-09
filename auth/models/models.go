package models

import (
	"time"
)

type User struct {
	Name      string    `db:"name" json:"name"`
	ID        int       `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"_"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
