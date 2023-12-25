package models

type Person struct {
	Id       int
	Name     string
	Email    string
	Surname  string
	Password string
	Role     int
}

type DbPerson struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Surname  string `db:"surname"`
	Password string `db:"password"`
	Role     int    `db:"role"`
}
