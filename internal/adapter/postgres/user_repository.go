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

func (u *PostgresUserRepository) CreateUser(email, name, password, role string) (*entity.User, error) {
	user := &entity.User{}
	query := `INSERT INTO account (email, name, password, role) VALUES ($1, $2, $3, $4) RETURNING id, email, password`
	err := u.db.QueryRow(query, email, name, password, role).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *PostgresUserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, email, name, password, role FROM account WHERE email = $1`
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *PostgresUserRepository) DeleteUser(userId int) error {
	query := `DELETE FROM account WHERE id = $1`
	_, err := u.db.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u *PostgresUserRepository) Find(userId int) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, email, name, password, role FROM account WHERE id = $1`
	err := u.db.QueryRow(query, userId).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *PostgresUserRepository) ListUsers() ([]entity.User, error) {
	query := `SELECT id, email, name, password, role FROM account`
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *PostgresUserRepository) UpdateUser(*entity.User) (*entity.User, error) {
	user := &entity.User{}
	query := `UPDATE account SET name = $1, role = $2`
	err := u.db.QueryRow(query, user.Name, user.Role).Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}
