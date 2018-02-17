package models

import (
	"time"
)

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
	CreatedDate       time.Time
}
