package repo

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hsmtkk/sturdy-giggle/model"
	"github.com/jackc/pgx/v5"
)

type User struct {
	conn *pgx.Conn
}

func NewUser(conn *pgx.Conn) *User {
	return &User{conn}
}

func (u *User) NewUser(ctx context.Context, name string, email string, password string) (model.User, error) {
	result := model.User{}
	uuid := uuid.New().String()
	hashedPassword := sha256Hash(password)
	createdAt := time.Now()
	if err := u.conn.QueryRow(ctx, `INSERT INTO users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) RETURNING id`, uuid, name, email, hashedPassword, createdAt).Scan(&result.ID); err != nil {
		return result, fmt.Errorf("insert failed: %w", err)
	}
	result.UUID = uuid
	result.Name = name
	result.Email = email
	result.Password = password
	result.CreatedAt = createdAt
	return result, nil
}

func sha256Hash(plainText string) string {
	sum := sha256.Sum256([]byte(plainText))
	hashed := fmt.Sprintf("%x", sum)
	return hashed
}
