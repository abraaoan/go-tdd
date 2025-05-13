package postgres

import (
	"database/sql"

	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/repository"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (u *PostgresUserRepository) CreateUser(email string, password string) (*entity.User, error) {
	user := &entity.User{}
	query := `INSERT INTO account (email, password) VALUES ($1, $2) RETURNING id, email, password`
	err := u.db.QueryRow(query, email, password).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *PostgresUserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, email, password FROM account WHERE email = $1`
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
