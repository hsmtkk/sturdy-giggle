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

func (u *User) New(ctx context.Context, name string, email string, password string) (model.User, error) {
	result := model.User{}
	result.UUID = uuid.New().String()
	result.Name = name
	result.Email = email
	result.Password = sha256Hash(password)
	result.CreatedAt = time.Now()
	if err := u.conn.QueryRow(ctx, `INSERT INTO users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) RETURNING id`, result.UUID, name, email, result.Password, result.CreatedAt).Scan(&result.ID); err != nil {
		return result, fmt.Errorf("insert failed: %w", err)
	}
	return result, nil
}

func sha256Hash(plainText string) string {
	sum := sha256.Sum256([]byte(plainText))
	hashed := fmt.Sprintf("%x", sum)
	return hashed
}

func (u *User) Get(ctx context.Context, id int64) (model.User, error) {
	result := model.User{}
	if err := u.conn.QueryRow(ctx, "SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1", id).Scan(&result.ID, &result.UUID, &result.Name, &result.Email, &result.Password, &result.CreatedAt); err != nil {
		return result, fmt.Errorf("select failed: %w", err)
	}
	return result, nil
}

func (u *User) Update(ctx context.Context, user model.User) error {
	if _, err := u.conn.Exec(ctx, "UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID); err != nil {
		return fmt.Errorf("update failed: %w", err)
	}
	return nil
}

func (u *User) Delete(ctx context.Context, id int64) error {
	if _, err := u.conn.Exec(ctx, "DELETE FROM users WHERE id = $1", id); err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	return nil
}
