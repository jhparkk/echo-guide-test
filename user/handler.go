package user

import (
	"jhpark/echo-guide-test/db/userdb"
)

type Handler struct {
	db userdb.DBHandler
}

func NewHandler(db userdb.DBHandler) *Handler {
	return &Handler{
		db: db,
	}
}
