package user

import (
	"jhpark/echo-guide-test/db/userdb"
	"jhpark/echo-guide-test/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// saveUser godoc
//
//	@Summary		[Summary] Register a new user
//	@Description	[Description] Register a new user
//	@ID				save-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		userResponse	true	"User info for registration"
//	@Success		201		{object}	userResponse
//	@Failure		400		{object}	utils.Error
//	@Failure		404		{object}	utils.Error
//	@Failure		500		{object}	utils.Error
//	@Router			/user [post]
func (h *Handler) saveUser(c echo.Context) error {
	name := c.FormValue("username")
	email := c.FormValue("email")

	user := &userdb.User{
		Username: name,
		Email:    email,
	}
	err := h.db.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	u := &userResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	// for i := 0; i < 10; i++ {

	// 	log.Println(i, ":", utils.GoroutineID(), "  sleep")
	// 	time.Sleep(time.Second * 1)
	// }

	// u := new(userResponse)
	// u.User.Username = "jhpark"
	// u.User.Email = "jhpark@sinsiway.com"
	return c.JSON(http.StatusOK, u)
}

// getUser godoc
//
//	@Summary		[Summary] Get the user by ID
//	@Description	[Description] Get the user by ID
//	@ID				get-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		201	{object}	userResponse
//	@Failure		400	{object}	utils.Error
//	@Failure		404	{object}	utils.Error
//	@Failure		500	{object}	utils.Error
//	@Security		ApiKeyAuth
//	@Router			/user/:id [get]
func (h *Handler) getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := new(userResponse)

	user, err := h.db.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	u.Username = user.Username
	u.Email = user.Email

	return c.JSON(http.StatusOK, u)
}

// showUser godoc
//
//	@Summary		[Summary] Get the user by ID
//	@Description	[Description] Get the user by ID
//	@ID				get-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_json	body		string	"user information json"
//	@Success		201	{object}	userResponse
//	@Failure		400	{object}	utils.Error
//	@Failure		404	{object}	utils.Error
//	@Failure		500	{object}	utils.Error
//	@Security		ApiKeyAuth
//	@Router			/user/:id [get]
func (h *Handler) showUser(c echo.Context) error {
	u := new(userRequest)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	//return c.JSONPretty(http.StatusOK, u, "##")
	// return below :
	// {
	// 	##"username": "jhpark_json",
	// 	##"email": "jhpark_json@sinsiway.com",
	// 	##"token": "token_data2"
	// }

	return c.JSON(http.StatusOK, u)
}

// updateUser godoc
//
//	@Summary		[Summary] update the user
//	@Description	[Description] update the user
//	@ID				update-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	flase	"name search"
//	@Success		201	{object}	userResponse
//	@Failure		400	{object}	utils.Error
//	@Failure		404	{object}	utils.Error
//	@Failure		500	{object}	utils.Error
//	@Security		ApiKeyAuth
//	@Router			/user [put]
func (h *Handler) updateUser(c echo.Context) error {
	name := c.QueryParam("username")
	u := &userResponse{
		Username: name,
		Email:    "updated@sinsiway.com",
	}
	return c.JSON(http.StatusOK, u)
}

// getUser godoc
//
//	@Summary		[Summary] delete the user
//	@Description	[Description] delete the user
//	@ID				delete-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	userResponse
//	@Failure		400	{object}	utils.Error
//	@Failure		404	{object}	utils.Error
//	@Failure		500	{object}	utils.Error
//	@Security		ApiKeyAuth
//	@Router			/user/:id [delete]
func (h *Handler) deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := new(userResponse)
	if id > 10 {
		u.Username = "deleted10over"
		u.Email = "deleted10over@sinsiway.com"
	} else if id == 10 {
		u.Username = "deleted10_"
		u.Email = "deleted10_@sinsiway.com"
	} else {
		u.Username = "deleted10under"
		u.Email = "deleted10under@sinsiway.com"
	}

	return c.JSON(http.StatusOK, u)
}
