package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/hsmtkk/sturdy-giggle/model"
	"github.com/jackc/pgx/v5"
)

type Todo struct {
	conn *pgx.Conn
}

func NewTodo(conn *pgx.Conn) *Todo {
	return &Todo{conn}
}

func (t *Todo) New(ctx context.Context, userID int64, content string) (model.Todo, error) {
	result := model.Todo{}
	result.UserID = userID
	result.Content = content
	result.CreatedAt = time.Now()
	if err := t.conn.QueryRow(ctx, `INSERT INTO users (userID, content, created_at) values ($1, $2, $3) RETURNING id`, userID, content, result.CreatedAt).Scan(&result.ID); err != nil {
		return result, fmt.Errorf("insert failed: %w", err)
	}
	return result, nil
}

func (t *Todo) Get(ctx context.Context, id int64) (model.Todo, error) {
	result := model.Todo{}
	if err := t.conn.QueryRow(ctx, "SELECT id, user_id, content, created_at FROM todos WHERE id = $1", id).Scan(&result.ID, &result.UserID, &result.Content, &result.CreatedAt); err != nil {
		return result, fmt.Errorf("select failed: %w", err)
	}
	return result, nil
}

func (t *Todo) List(ctx context.Context) ([]model.Todo, error) {
	results := []model.Todo{}
	rows, err := t.conn.Query(ctx, "SELECT id, user_id, content, created_at FROM todos")
	if err != nil {
		return nil, fmt.Errorf("select failed: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		result := model.Todo{}
		if err := rows.Scan(&result.ID, &result.UserID, &result.Content, &result.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func (t *Todo) Update(ctx context.Context, todo model.Todo) error {
	if _, err := t.conn.Exec(ctx, "UPDATE todos SET content = $1 WHERE id = $2", todo.Content, todo.ID); err != nil {
		return fmt.Errorf("update failed: %w", err)
	}
	return nil
}

func (t *Todo) Delete(ctx context.Context, id int64) error {
	if _, err := t.conn.Exec(ctx, "DELETE FROM todos WHERE id = $1", id); err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	return nil
}
