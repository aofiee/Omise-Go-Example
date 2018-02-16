package tests

import (
	"net/url"
	"time"

	"github.com/aofiee666/OmiseWallet/app/models"
	"github.com/revel/revel/testing"
)

// AppTest struct
type AppTest struct {
	testing.TestSuite
}

// Before func reciever
func (t *AppTest) Before() {
	println("Set up")
}

// TestThatIndexPageWorks func reciever
func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// After func reciever
func (t *AppTest) After() {
	println("Tear down")
}

// TestUserModel func reciever
func (t *AppTest) TestUserModel() {
	ts := time.Now()
	user := models.User{
		Username:    "test",
		Password:    "password",
		Role:        1,
		CreatedDate: ts,
	}
	t.AssertEqual(user.Username, "test")
	t.AssertEqual(user.Password, "password")
	t.AssertEqual(user.Role, 1)
	t.AssertEqual(user.CreatedDate, ts)
}

// TestOmiseKeyModel func reciever
func (t *AppTest) TestOmiseKeyModel() {
	ts := time.Now()
	omise := models.OmiseKey{
		PublicKey:   "PublicKey",
		SecretKey:   "SecretKey",
		CreatedDate: ts,
	}
	t.AssertEqual(omise.PublicKey, "PublicKey")
	t.AssertEqual(omise.SecretKey, "SecretKey")
	t.AssertEqual(omise.CreatedDate, ts)
}

// TestThatLoginPageWorks func reciever
func (t *AppTest) TestThatLoginPageWorks() {
	t.PostForm("/Login", url.Values{"username": {"test"}, "password": {"test"}, "remember": {"on"}})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// TestThatDashboardPageWorks func reciever
func (t *AppTest) TestThatDashboardPageWorks() {
	t.Get("/Dashboard")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
