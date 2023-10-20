package model

import "time"

type User struct {
	ID        int64
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
