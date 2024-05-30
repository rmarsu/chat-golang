package repository

import (
	chat "chat"
	"context"
	"database/sql"
	// "chat/pkg/service"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) chat.Repository {
	return &repository{db: db}

}

func (r *repository) CreateUser(ctx context.Context, user *chat.User) (*chat.User, error) {
	var lastInsertID int64
	query := "INSERT INTO users(username, email , password) VALUES ($1, $2 , $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertID)
	if err != nil {
		return &chat.User{}, err
	}
	user.ID = int64(lastInsertID)
	return user, nil

}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*chat.User, error) {
	u := chat.User{}

	query := "SELECT id, email, username , password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &chat.User{}, err
	}
	return &u, nil
}
