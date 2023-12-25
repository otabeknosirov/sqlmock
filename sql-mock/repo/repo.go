package repo

import (
	"database/sql"
	"tests/sql-mock/models"
)

type Repo interface {
	Close()
	Begin(db *sql.DB) (*sql.Tx, error)
	Create(user *models.Person) error
	Update(user *models.Person) error
	Find() ([]*models.Person, error)
	FindById(id int) (models.Person, error)
	Delete(id int) error
}
