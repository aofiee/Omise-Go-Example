package models

import (
	"time"
)

//Recipient struct
type Recipient struct {
	ID                int16
	RecipientName     string
	Description       string
	Email             string
	RecipientType     string
	TaxID             string
	BankAccountBrand  string
	BankAccountNumber string
	BankAccountName   string
	IsDefault         int
	OmiseID           string
	CreatedDate       time.Time
}
