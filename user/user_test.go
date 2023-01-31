package user

import (
	"encoding/json"
	"fmt"
	"jhpark/echo-guide-test/db/userdb"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	setUp()
	defer tearDown()

	assert := assert.New(t)
	//req := httptest.NewRequest(echo.POST, "/api/user?username=test_user&email=test_user@sinsiway.com", nil)
	// set quert params
	q := make(url.Values)
	q.Set("username", "test_user")
	q.Set("email", "test_user@sinsiway.com")
	req := httptest.NewRequest(echo.POST, "/?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// assertions
	if assert.NoError(h.saveUser(c)) {
		fmt.Println("response body : ", string(rec.Body.Bytes()))
		u := &userResponse{}
		err := json.Unmarshal(rec.Body.Bytes(), u)
		assert.NoError(err)
		assert.Equal("test_user", u.Username)
		assert.Equal("test_user@sinsiway.com", u.Email)
	}
}

func TestGetUser(t *testing.T) {
	setUp()
	defer tearDown()

	assert := assert.New(t)
	user := &userdb.User{
		Username: "test_user",
		Email:    "test_user@sinsiway.com",
	}

	err := testdb.CreateUser(user)
	assert.NoError(err)
	fmt.Println("the user created : ", user)

	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/profiles/:username/follow")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(user.ID)))

	// assertions
	if assert.NoError(h.getUser(c)) {
		fmt.Println("response body : ", string(rec.Body.Bytes()))
		u := &userResponse{}
		err := json.Unmarshal(rec.Body.Bytes(), u)
		assert.NoError(err)
		assert.Equal("test_user", u.Username)
		assert.Equal("test_user@sinsiway.com", u.Email)
	}
}

func TestSampleC(t *testing.T) {
	setUp()
	defer tearDown()
	func() {
		fmt.Println("TestSampleC !!!")
	}()
}
