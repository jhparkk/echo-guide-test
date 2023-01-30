package userdb

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqlite3Handler struct {
	db *gorm.DB
}

func newSqlite3Handler(filepath string) DBHandler {
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	user := &User{}
	err = db.AutoMigrate(user)
	if err != nil {
		panic(err)
	}
	h := &sqlite3Handler{
		db: db,
	}
	return h
}

func deleteSqlite3DbFile(filepath string) error {
	if err := os.Remove("./../user_test.db"); err != nil {
		return err
	}
	return nil
}

func (h *sqlite3Handler) Close() {
	db, _ := h.db.DB()
	db.Close()
}

func (h *sqlite3Handler) GetUserByID(id int) (*User, error) {
	var user User
	err := h.db.Table("users").Where(id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (h *sqlite3Handler) GetUserByName(username string) ([]*User, error) {
	var users []*User
	err := h.db.Table("users").Where("username = ?", username).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (h *sqlite3Handler) CreateUser(u *User) error {
	return h.db.Table("users").Create(u).Error
}

func (h *sqlite3Handler) UpdateUser(u *User) error {
	return h.db.Table("users").Updates(u).Error
}

func (h *sqlite3Handler) DeleteUser(id int) error {
	return h.db.Table("users").Delete(&User{}, id).Error
}
