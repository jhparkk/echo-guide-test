package user

import (
	"fmt"
	"jhpark/echo-guide-test/db/userdb"
	"jhpark/echo-guide-test/router"
	"log"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

var (
	isSetUpTestCase bool
	testdb          userdb.DBHandler
	h               *Handler
	e               *echo.Echo
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain!! Before the tests")
	setUpTestCase()

	code := m.Run()

	tearDownTestCase()
	fmt.Println("TestMain!! After the tests")
	os.Exit(code)
}

func setUpTestCase() {
	testdb = userdb.NewTestDBHandler()
	h = NewHandler(testdb)
	e = router.New()
	isSetUpTestCase = true
}

func tearDownTestCase() {
	testdb.Close()

	if err := userdb.DropTestDBHandler(); err != nil {
		log.Fatal(err)
	}

	isSetUpTestCase = false
}

func setUp() {
	fmt.Println("@setUp")
	testdb.DeleteAllUsers()
}

func tearDown() {
	fmt.Println("@tearDown")
}
