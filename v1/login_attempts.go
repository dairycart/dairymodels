package models

import (
	"time"
)

// LoginAttempt represents a Dairycart login attempt
type LoginAttempt struct {
	Username   string    `json:"username"`   // username
	Successful bool      `json:"successful"` // successful
	ID         uint64    `json:"id"`         // id
	CreatedOn  time.Time `json:"created_on"` // created_on
}

type LoginAttemptListResponse struct {
	ListResponse
	LoginAttempts []LoginAttempt `json:"login_attempts"`
}
