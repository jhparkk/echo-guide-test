package userdb

type DBHandler interface {
	GetUserByID(id int) (*User, error)
	GetUserByName(username string) ([]*User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(id int) error
}

func NewDBHandler() DBHandler {
	return newSqlite3Handler("./user.db")
}

func NewTest() DBHandler {
	return newSqlite3Handler("./../user_test.db")
}

func DropTest() error {
	return deleteSqlite3DbFile("./../user_test.db")
}
