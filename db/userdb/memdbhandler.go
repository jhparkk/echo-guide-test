package userdb

import (
	"fmt"
	"time"

	"golang.org/x/exp/maps"
)

type memdbHandler struct {
	users map[int]*User
}

func newMemDBHandler() DBHandler {
	m := &memdbHandler{}
	m.users = make(map[int]*User)
	return m
}

func (h *memdbHandler) Close() {

}

func (h *memdbHandler) GetUserByID(id int) (*User, error) {
	user := h.users[id]
	return user, nil
}

func (h *memdbHandler) GetUserByName(username string) ([]*User, error) {
	var users []*User
	for _, u := range h.users {
		if u.Username == username {
			users = append(users, u)
		}
	}
	return users, nil
}

func (h *memdbHandler) CreateUser(u *User) error {
	id := len(h.users) + 1
	h.users[id] = u
	return nil
}

func (h *memdbHandler) UpdateUser(u *User) error {
	user := h.users[int(u.ID)]
	if user == nil {
		return fmt.Errorf("user not found : %v\n", u)
	}
	user.UpdatedAt = time.Now()
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password

	return nil
}

func (h *memdbHandler) DeleteUser(id int) error {
	if _, ok := h.users[id]; ok {
		delete(h.users, id)
	}

	return fmt.Errorf("user[id:%d] not found.\n", id)
}

func (h *memdbHandler) DeleteAllUsers() error {
	maps.Clear(h.users)
	return nil
}
