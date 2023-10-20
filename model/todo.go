package model

import "time"

type Todo struct {
	ID        int64
	UserID    int64
	Content   string
	CreatedAt time.Time
}
