package models

import (
	"time"
)

// PasswordResetToken represents a Dairycart password reset token
type PasswordResetToken struct {
	PasswordResetOn NullTime  `json:"password_reset_on"` // password_reset_on
	Token           string    `json:"token"`             // token
	ID              uint64    `json:"id"`                // id
	CreatedOn       time.Time `json:"created_on"`        // created_on
	UserID          uint64    `json:"user_id"`           // user_id
	ExpiresOn       time.Time `json:"expires_on"`        // expires_on
}
