package models

import (
	"time"
)

// OmiseKey Struct
type OmiseKey struct {
	ID          int
	PublicKey   string
	SecretKey   string
	CreatedDate time.Time
}
