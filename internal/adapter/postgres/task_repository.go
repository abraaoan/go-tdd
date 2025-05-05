package postgres

import (
	"database/sql"

	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/repository"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) repository.TaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (t *PostgresTaskRepository) Delete(id int, userID int) error {
	query := `DELETE FROM task WHERE id = $1 and user_id = $2`
	_, err := t.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	return nil
}

func (t *PostgresTaskRepository) FindById(id int) (*entity.Task, error) {
	var task entity.Task
	query := `SELECT id, title, status, user_id FROM task WHERE id = $1`
	err := t.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Status, &task.UserID)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *PostgresTaskRepository) List(userID int) ([]entity.Task, error) {
	query := `SELECT id, title, status, user_id FROM task WHERE user_id = $1 ORDER BY title`
	rows, err := t.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.UserID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *PostgresTaskRepository) Save(title string, userID int) (*entity.Task, error) {
	task := &entity.Task{}
	query := `INSERT INTO task (title, status, user_id) VALUES ($1, $2, $3) RETURNING id, title, status, user_id`
	err := t.db.QueryRow(query, title, false, userID).Scan(&task.ID, &task.Title, &task.Status, &task.UserID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *PostgresTaskRepository) Update(task *entity.Task) error {
	query := "UPDATE task SET title = $1, status = $2 WHERE id = $3"
	_, err := t.db.Exec(query, task.Title, task.Status, task.ID)
	if err != nil {
		return err
	}

	return nil
}
