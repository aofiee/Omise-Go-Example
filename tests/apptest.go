package tests

import (
	"net/url"
	"time"

	"github.com/aofiee666/OmiseWallet/app/controllers"
	"github.com/aofiee666/OmiseWallet/app/models"
	"github.com/revel/revel/testing"
)

// AppTest struct
type AppTest struct {
	testing.TestSuite
	controllers.App
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

// TestRecipientModel func reciever
func (t *AppTest) TestRecipientModel() {
	ts := time.Now()
	recip := models.Recipient{
		RecipientName:     "Khomkrid Lerdprasert",
		RecipientType:     "Individual",
		Description:       "Test",
		Email:             "aofiee666@gmail.com",
		TaxID:             "1234567890",
		BankAccountBrand:  "KBank",
		BankAccountName:   "Khomkrid Lerdprasert",
		BankAccountNumber: "1234567890",
		IsDefault:         1,
		OmiseID:           "text",
		CreatedDate:       ts,
	}
	t.AssertEqual(recip.RecipientName, "Khomkrid Lerdprasert")
	t.AssertEqual(recip.RecipientType, "Individual")
	t.AssertEqual(recip.Description, "Test")
	t.AssertEqual(recip.Email, "aofiee666@gmail.com")
	t.AssertEqual(recip.TaxID, "1234567890")
	t.AssertEqual(recip.BankAccountBrand, "KBank")
	t.AssertEqual(recip.BankAccountName, "Khomkrid Lerdprasert")
	t.AssertEqual(recip.BankAccountNumber, "1234567890")
	t.AssertEqual(recip.IsDefault, 1)
	t.AssertEqual(recip.OmiseID, "text")
	t.AssertEqual(recip.CreatedDate, ts)
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

//TestPublickeyPageWorks func
func (t *AppTest) TestPublickeyPageWorks() {
	t.Get("/Publickey")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// TestUpdateKeyPageWorks func reciever
func (t *AppTest) TestUpdateKeyPageWorks() {
	t.PostForm("/UpdateKey", url.Values{"publickey": {"test"}, "secretkey": {"test"}})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

//TestBankDefaultPageWorks func
func (t *AppTest) TestBankDefaultPageWorks() {
	t.Get("/DefaultBank")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

//TestUpdateBankDefaultPageWorks func
func (t *AppTest) TestUpdateBankDefaultPageWorks() {
	t.PostForm("/UpdateDefaultBank", url.Values{"optradio": {"test"}, "name": {"test"}, "email": {"test"}, "taxid": {"test"}, "description": {"test"}, "bankaccountbrand": {"test"}, "bankaccountname": {"test"}, "bankaccountnumber": {"test"}})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
