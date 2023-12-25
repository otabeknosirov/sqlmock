package mock

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"tests/sql-mock/models"
	"tests/sql-mock/repo"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("can't connect to db")
	}
	defer func() {
		repo.Close()
	}()
	user := &models.DbPerson{
		Id:    1,
		Name:  "Otabek",
		Email: "otabek94_30@mail.ru",
	}
	q := `insert into users`
	mock.ExpectPrepare(q).ExpectExec().WithArgs(user.Id, user.Name, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	u := &models.Person{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	err = repo.Create(u)
	assert.NoError(t, err)
	//fmt.Println(assert.NoError(t, err))
}

func TestCreateWithError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("can't connect to db")
	}
	defer func() {
		repo.Close()
	}()
	user := &models.DbPerson{
		Id:    1,
		Name:  "Otabek",
		Email: "otabek94_30@mail.ru",
	}
	q := `insert into users`
	mock.ExpectPrepare(q).ExpectExec().WithArgs(user.Id, user.Email).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.Create(&models.Person{
		Id:    1,
		Email: "otabek94_30@mail.ru",
	})
	//fmt.Println(err, e)
	assert.Error(t, err)
}

func TestUpdateNoError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}
	defer func() {
		repo.Close()
	}()

	mUser := models.DbPerson{
		Id:    1,
		Name:  "Otaiba",
		Email: "otaiba94_30@mail.ru",
	}
	q := `update users`
	mock.ExpectPrepare(q).ExpectExec().WithArgs(mUser.Name, mUser.Email, mUser.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Update(&models.Person{Id: 1, Name: "Otaiba", Email: "otaiba94_30@mail.ru"})
	assert.NoError(t, err)
}

func TestUpdateWithError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}

	defer func() {
		repo.Close()
	}()
	mUser := models.DbPerson{
		Id:    1,
		Name:  "Otaiba",
		Email: "otaiba94_30@mail.ru",
	}
	q := `update users`
	mock.ExpectPrepare(q).ExpectExec().WithArgs(mUser.Name, mUser.Id).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.Update(&models.Person{
		Id:   1,
		Name: "Otaiba",
	})
	assert.Error(t, err)
}

func TestUpdateNoSqlError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}

	defer func() {
		repo.Close()
	}()
	mUser := models.DbPerson{
		Id:    2,
		Name:  "Otaiba",
		Email: "otaiba94_30@mail.ru",
	}
	q := `update users`
	mock.ExpectPrepare(q).ExpectExec().WithArgs(mUser.Name, mUser.Id).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.Update(&models.Person{
		Id:   2,
		Name: "Otaiba",
	})
	assert.Error(t, err)
}

func TestDeleteNoError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}

	defer func() {
		repo.Close()
	}()
	q := `DELETE FROM users WHERE`
	mUser := models.DbPerson{Id: 1}
	mock.ExpectPrepare(q).ExpectExec().WithArgs(mUser.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Delete(models.DbPerson{Id: 1}.Id)
	assert.NoError(t, err)
}

func TestDeleteWithError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}

	defer func() {
		repo.Close()
	}()
	q := `DELETE FROM users WHERE`
	mUser := models.DbPerson{Id: 1}
	mock.ExpectPrepare(q).ExpectExec().WithArgs(mUser).WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Delete(models.DbPerson{Id: 1}.Id)
	assert.Error(t, err)
}

func TestGetById(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}

	defer func() {
		repo.Close()
	}()
	mUser := models.DbPerson{Id: 1, Name: "Otabek", Email: "otabek94_30@mail.ru"}
	q := `select id,name,email from users where`
	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(mUser.Id, mUser.Name, mUser.Email)
	mock.ExpectQuery(q).WithArgs(mUser.Id).WillReturnRows(rows)

	user, err := repo.FindById(models.DbPerson{Id: 1}.Id)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
func TestGetByIdWithError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}
	defer func() {
		repo.Close()
	}()

	mUser := models.DbPerson{Id: 2}
	q := `select id,name,email from users where`

	rows := sqlmock.NewRows([]string{"id", "name", "email"})
	mock.ExpectQuery(q).WithArgs(mUser.Id).WillReturnRows(rows)

	user, err := repo.FindById(models.DbPerson{Id: 2}.Id)
	fmt.Println(err)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestGetByIdTypeError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}
	defer func() {
		repo.Close()
	}()

	mUser := models.DbPerson{Id: 1}
	q := `select id,name,email from users where`

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, nil, "otabek94_30@mail.ru")
	mock.ExpectQuery(q).WithArgs(mUser.Id).WillReturnRows(rows)

	_, err = repo.FindById(models.DbPerson{Id: 1}.Id)
	assert.EqualError(t, err, "sql: Scan error on column index 1, name \"name\": converting NULL to string is unsupported")
	assert.Error(t, err)
}

func TestGetAllUsers(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}
	defer func() {
		repo.Close()
	}()

	rowData := [][]driver.Value{
		{
			1, "Otabek", "otabek94_30@mail.ru",
		},
		{
			2, "Ilyosbek", "ilyosbek@mail.ru",
		},
	}
	q := `select id,name,email from users`

	rows := mock.NewRows([]string{"id", "name", "email"}).AddRows(rowData...)
	mock.ExpectQuery(q).WillReturnRows(rows)

	users, err := repo.Find()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
}

func TestGetAllUsersWithError(t *testing.T) {
	db, mock := NewMock()
	repo, err := repo.NewRepository(db)
	if err != nil {
		t.Fatalf("%v cannot connect to Db", err)
	}
	defer func() {
		repo.Close()
	}()

	q := `select id,name,email from users`

	rows := mock.NewRows([]string{"id", "name", "email"})
	mock.ExpectQuery(q).WillReturnRows(rows).WillReturnError(sql.ErrNoRows)

	users, err := repo.Find()
	assert.Empty(t, users)
	assert.NotNil(t, err)
	assert.Error(t, err)
}
