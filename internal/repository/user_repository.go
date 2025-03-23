package repository

import (
	"context"

	"github.com/Numbone/user-api/internal/model"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	row := r.db.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id = $1", id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.Exec(ctx, "UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, user.ID)
	return err
}
