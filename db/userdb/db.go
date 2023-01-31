package userdb

type DBHandler interface {
	Close()
	GetUserByID(id int) (*User, error)
	GetUserByName(username string) ([]*User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(id int) error
	DeleteAllUsers() error
}

func NewDBHandler() DBHandler {
	return newSqlite3Handler("./user.db")
}

func NewTestDBHandler() DBHandler {
	return newSqlite3Handler("./../user_test.db")
}

func DropTestDBHandler() error {
	return deleteSqlite3DbFile("./../user_test.db")
}
