package repo

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"tests/sql-mock/models"
	"time"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) (Repo, error) {
	rp := Repository{
		Db: db,
	}
	return &rp, nil
}

func ConnectToDb(dns string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (r *Repository) Close() {
	r.Db.Close()
}
func (r *Repository) Begin(db *sql.DB) (*sql.Tx, error) {

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r *Repository) Create(user *models.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	q := `insert into users(id,name,email) values($1,$2,$3)`
	stmt, err := r.Db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	if _, err = stmt.ExecContext(ctx, user.Id, user.Name, user.Email); err != nil {
		return err
	}
	return nil
}
func (r *Repository) Update(user *models.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	q := "update users set name=$1, email=$2 where id=$3"
	stmt, err := r.Db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Email, user.Id)
	return err
}

func (r *Repository) Find() ([]*models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []*models.Person
	q := `select id,name,email from users`
	rows, err := r.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.DbPerson
		err = rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &models.Person{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return users, nil
}

func (r *Repository) FindById(id int) (models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userDb := models.DbPerson{}
	row := r.Db.QueryRowContext(ctx, "select id,name,email from users where id=$1", id)
	if err := row.Scan(&userDb.Id, &userDb.Name, &userDb.Email); err != nil {
		return models.Person{}, err
	}
	return models.Person{
		Id:    userDb.Id,
		Name:  userDb.Name,
		Email: userDb.Email,
	}, nil
}

func (r *Repository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"
	stmt, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
