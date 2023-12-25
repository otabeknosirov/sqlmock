package main

import (
	"flag"
	"fmt"
	"tests/sql-mock/models"
	"tests/sql-mock/repo"
)

type App struct {
	DSN  string
	repo *repo.Repo
}

func main() {
	app := App{}
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable timezone=UTC connect_timeout=5", "Posgtres connection")
	db, err := repo.ConnectToDb(app.DSN)
	if err != nil {
		panic(err)
	}
	rp, err := repo.NewRepository(db)
	//tx, err := rp.Begin(rp.Db)
	//if err != nil{
	//	panic(err)
	//}
	app.repo = &rp

	err = rp.Create(&models.Person{
		Id:    1,
		Name:  "Otabek",
		Email: "otabek94_30@mail.ru",
	})
	if err != nil {
		fmt.Println(rp, err)
		panic(err)
	}

	_, err = rp.Find()
	if err != nil {
		fmt.Println(rp, err)
		panic(err)
	}

}
