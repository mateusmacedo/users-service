package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mateusmacedo/users-service/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(dbFile string, tableName string) (*SQLiteUserRepository, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL
	)`, tableName)
	_, err = db.Exec(query)

	if err != nil {
		return nil, err
	}

	return &SQLiteUserRepository{db: db}, nil
}

func (r *SQLiteUserRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (id, name) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *SQLiteUserRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, name FROM users WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *SQLiteUserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *SQLiteUserRepository) List(ctx context.Context, filter map[string]interface{}, limit int, offset int) ([]*domain.User, error) {
	query := `SELECT id, name FROM users LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
