package user

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(v1 *echo.Group) {
	v1.POST("/user", h.saveUser)   // form application/x-www-form-urlencoded
	v1.GET("/user/:id", h.getUser) // path params
	v1.GET("/user", h.showUser)    // json params
	v1.PUT("/user", h.updateUser)  // query params
	v1.DELETE("/user/:id", h.deleteUser)
}
