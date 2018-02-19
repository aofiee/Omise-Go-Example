package controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aofiee/Omise-Go-Example/app/models"
	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"

	"github.com/revel/revel"
)

// Dashboard structure
type Dashboard struct {
	*revel.Controller
	App
}

var (
	myName string
)

// Index method
func (c Dashboard) Index() revel.Result {
	myName := strings.Title(c.Session["username"])

	var d Dashboard
	p, s := d.getPublicAndSecretKey()
	fmt.Println(p, s)
	return c.Render(myName)
}

//checkUser func
func (c Dashboard) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in before")
		return c.Redirect(App.Index)
	}
	return nil
}

//Logout func
func (c Dashboard) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	c.ViewArgs["username"] = nil
	return c.Redirect(App.Index)
}

// PublicKey func
func (c Dashboard) PublicKey() revel.Result {
	// myName := strings.Title(c.Session["username"])
	db := models.Gorm
	var omise models.OmiseKey
	db.First(&omise)
	publickey := omise.PublicKey
	secretkey := omise.SecretKey
	return c.Render(myName, publickey, secretkey)
}

// UpdateKey func
func (c Dashboard) UpdateKey(publickey string, secretkey string) revel.Result {
	c.Validation.Required(publickey)
	c.Validation.Required(secretkey)
	if c.Validation.HasErrors() {
		c.Flash.Error("กรุณากรอก public key และ secret key ด้วยครับ")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Dashboard.PublicKey)
	}
	myName := strings.Title(c.Session["username"])
	db := models.Gorm
	var omise models.OmiseKey
	db.First(&omise)

	if omise.ID == 0 {
		db.FirstOrCreate(&omise, models.OmiseKey{
			PublicKey:   publickey,
			SecretKey:   secretkey,
			CreatedDate: time.Now(),
		})
	} else {
		omise.PublicKey = publickey
		omise.SecretKey = secretkey
		db.Save(&omise)
	}
	c.ViewArgs["myName"] = myName
	c.ViewArgs["publickey"] = publickey
	c.ViewArgs["secretkey"] = secretkey
	return c.RenderTemplate("Dashboard/PublicKey.html")
}

//DefaultBank func
func (c Dashboard) DefaultBank() revel.Result {
	myName := strings.Title(c.Session["username"])
	db := models.Gorm
	var recipient models.Recipient
	db.Where("is_default = 1").First(&recipient)
	actionURL := "UpdateDefaultBank"
	return c.Render(myName, recipient, actionURL)
}

//UpdateDefaultBank func
func (c Dashboard) UpdateDefaultBank(optradio string, name string, email string, taxid string, description string, bankaccountbrand string, bankaccountname string, bankaccountnumber string) revel.Result {
	myName := strings.Title(c.Session["username"])
	c.ViewArgs["myName"] = myName
	c.Validation.Required(optradio)
	c.Validation.Required(email)
	c.Validation.Required(name)
	c.Validation.Required(taxid)
	c.Validation.Required(description)
	c.Validation.Required(bankaccountbrand)
	c.Validation.Required(bankaccountname)
	c.Validation.Required(bankaccountnumber)
	if c.Validation.HasErrors() {
		c.Flash.Error("กรุณากรอกข้อมูลให้ครบด้วยครับ")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Dashboard.DefaultBank)
	}
	db := models.Gorm
	var recipient models.Recipient
	db.Where("is_default = 1").First(&recipient)
	if recipient.ID == 0 {
		db.FirstOrCreate(&recipient, models.Recipient{
			RecipientName:     name,
			Description:       description,
			Email:             email,
			RecipientType:     optradio,
			TaxID:             taxid,
			BankAccountBrand:  bankaccountbrand,
			BankAccountName:   bankaccountname,
			BankAccountNumber: bankaccountnumber,
			IsDefault:         1,
			CreatedDate:       time.Now(),
		})
		/*
		   +----------------------------------------------------------------------+
		   |                integrate with  omise recipient api                   |
		   |                    return omise's recipient key                      |
		   +----------------------------------------------------------------------+
		*/
		go recipientSaveInOmise(recipient)
		/*
		   +----------------------------------------------------------------------+
		   |                integrate with  omise recipient api                   |
		   |                    return omise's recipient key                      |
		   +----------------------------------------------------------------------+
		*/
	} else {
		recipient.RecipientName = name
		recipient.Description = description
		recipient.Email = email
		recipient.RecipientType = optradio
		recipient.TaxID = taxid
		recipient.BankAccountBrand = bankaccountbrand
		recipient.BankAccountName = bankaccountname
		recipient.BankAccountNumber = bankaccountnumber
		recipient.IsDefault = 1
		recipient.CreatedDate = time.Now()
		db.Save(&recipient)
		/*
		   +----------------------------------------------------------------------+
		   |                integrate with  omise recipient api                   |
		   |                    return omise's recipient key                      |
		   +----------------------------------------------------------------------+
		*/
		go recipientUpdateInOmise(recipient)
		/*
		   +----------------------------------------------------------------------+
		   |                integrate with  omise recipient api                   |
		   |                    return omise's recipient key                      |
		   +----------------------------------------------------------------------+
		*/
	}
	c.Flash.Success("Completed...")
	c.ViewArgs["recipient"] = recipient
	c.ViewArgs["actionURL"] = "UpdateDefaultBank"
	return c.RenderTemplate("Dashboard/DefaultBank.html")
}

//recipientUpdateInOmise func for run go
func recipientUpdateInOmise(recipient models.Recipient) bool {
	var d Dashboard
	OmisePublicKey, OmiseSecretKey := d.getPublicAndSecretKey()
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
		return false
	}
	var typeBank omise.RecipientType
	if recipient.RecipientType == "individual" {
		typeBank = omise.Individual
	} else {
		typeBank = omise.Corporation
	}
	omiseRecipient, updateRecipient := &omise.Recipient{}, &operations.UpdateRecipient{
		RecipientID: recipient.OmiseID,
		Name:        recipient.RecipientName,
		Email:       recipient.Email,
		Description: recipient.Description,
		Type:        typeBank,
		TaxID:       recipient.TaxID,
		BankAccount: &omise.BankAccount{
			Brand:  recipient.BankAccountBrand,
			Number: recipient.BankAccountNumber,
			Name:   recipient.BankAccountName,
		},
	}
	if e := client.Do(omiseRecipient, updateRecipient); e != nil {
		log.Fatal(e)
		return false
	}
	db := models.Gorm
	var recipientDB models.Recipient
	db.Where("id = ?", recipient.ID).First(&recipientDB)
	recipientDB.OmiseID = omiseRecipient.ID
	db.Save(&recipient)
	return true
}

//recipientSaveInOmise func for run go
func recipientSaveInOmise(recipient models.Recipient) bool {
	var d Dashboard
	OmisePublicKey, OmiseSecretKey := d.getPublicAndSecretKey()
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
		return false
	}
	var typeBank omise.RecipientType
	if recipient.RecipientType == "individual" {
		typeBank = omise.Individual
	} else {
		typeBank = omise.Corporation
	}

	omiseRecipient, createRecipient := &omise.Recipient{}, &operations.CreateRecipient{
		Name:        recipient.RecipientName,
		Email:       recipient.Email,
		Description: recipient.Description,
		TaxID:       recipient.TaxID,
		Type:        typeBank,
		BankAccount: &omise.BankAccount{
			Brand:  recipient.BankAccountBrand,
			Number: recipient.BankAccountNumber,
			Name:   recipient.BankAccountName,
		},
	}
	if e := client.Do(omiseRecipient, createRecipient); e != nil {
		log.Fatal(e)
		return false
	}
	db := models.Gorm
	var recipientDB models.Recipient
	db.Where("id = ?", recipient.ID).First(&recipientDB)
	recipientDB.OmiseID = omiseRecipient.ID
	db.Save(&recipientDB)
	fmt.Println("omise", omiseRecipient)
	return true
}

//ListAllRecipient func
func (c Dashboard) ListAllRecipient() revel.Result {
	myName := strings.Title(c.Session["username"])
	/*
	   +------------------------------------------------------------------------------+
	   |                                                                              |
	   |                 integrate with  omise list recipient api                     |
	   |                                                                              |
	   +------------------------------------------------------------------------------+
	*/
	recipients := listRecipient()
	/*
	   +------------------------------------------------------------------------------+
	   |                                                                              |
	   |                 integrate with  omise list recipient api                     |
	   |                                                                              |
	   +------------------------------------------------------------------------------+
	*/
	var sliceRecipient []*omise.Recipient

	for _, item := range recipients.Data {
		sliceRecipient = append(sliceRecipient, item)
	}
	fmt.Println(sliceRecipient)
	return c.Render(myName, sliceRecipient)
}

func listRecipient() (omiseRecipient omise.RecipientList) {
	/*
	   +------------------------------------------------------------------------------+
	   |                                                                              |
	   |                 integrate with  omise list recipient api                     |
	   |                                                                              |
	   +------------------------------------------------------------------------------+
	*/
	var d Dashboard
	OmisePublicKey, OmiseSecretKey := d.getPublicAndSecretKey()
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
		return omiseRecipient
	}

	recipients, listRecipients := &omise.RecipientList{}, &operations.ListRecipients{
		List: operations.List{
			Offset: 0,
			Limit:  10,
		},
	}
	if e := client.Do(recipients, listRecipients); e != nil {
		log.Fatal(e)
		return omiseRecipient
	}
	omiseRecipient = *recipients
	/*
	   +------------------------------------------------------------------------------+
	   |                                                                              |
	   |                 integrate with  omise list recipient api                     |
	   |                                                                              |
	   +------------------------------------------------------------------------------+
	*/
	return omiseRecipient
}

//NewRecipientForm func
func (c Dashboard) NewRecipientForm() revel.Result {
	myName := strings.Title(c.Session["username"])
	c.ViewArgs["myName"] = myName
	c.ViewArgs["actionURL"] = "SaveNewRecipient"
	return c.RenderTemplate("Dashboard/DefaultBank.html")
}

//SaveNewRecipient func
func (c Dashboard) SaveNewRecipient(optradio string, name string, email string, taxid string, description string, bankaccountbrand string, bankaccountname string, bankaccountnumber string) revel.Result {
	myName := strings.Title(c.Session["username"])
	/*
	   +------------------------------------------------------------------------------+
	   |                                                                              |
	   |                             Validate Form Data                               |
	   |                                                                              |
	   +------------------------------------------------------------------------------+
	*/
	c.Validation.Required(optradio)
	c.Validation.Required(email)
	c.Validation.Required(name)
	c.Validation.Required(taxid)
	c.Validation.Required(description)
	c.Validation.Required(bankaccountbrand)
	c.Validation.Required(bankaccountname)
	c.Validation.Required(bankaccountnumber)
	if c.Validation.HasErrors() {
		c.Flash.Error("กรุณากรอกข้อมูลให้ครบด้วยครับ")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Dashboard.DefaultBank)
	}
	/*
	   +------------------------------------------------------------------------------+
	   |                                                                              |
	   |                             Validate Form Data                               |
	   |                                                                              |
	   +------------------------------------------------------------------------------+
	*/
	db := models.Gorm
	var recipient models.Recipient
	recipient.RecipientName = name
	recipient.Description = description
	recipient.Email = email
	recipient.RecipientType = optradio
	recipient.TaxID = taxid
	recipient.BankAccountBrand = bankaccountbrand
	recipient.BankAccountName = bankaccountname
	recipient.BankAccountNumber = bankaccountnumber
	recipient.IsDefault = 0
	recipient.CreatedDate = time.Now()
	db.Save(&recipient)
	/*
		+----------------------------------------------------------------------+
		|                integrate with  omise recipient api                   |
		|                    return omise's recipient key                      |
		+----------------------------------------------------------------------+
	*/
	go recipientSaveInOmise(recipient)
	/*
		+----------------------------------------------------------------------+
		|                integrate with  omise recipient api                   |
		|                    return omise's recipient key                      |
		+----------------------------------------------------------------------+
	*/
	c.ViewArgs["myName"] = myName
	c.ViewArgs["actionURL"] = "SaveNewRecipient"
	return c.Redirect("/ListAllRecipient")
}
