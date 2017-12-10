package models

import (
	"time"
)

// User represents a Dairycart user
type User struct {
	Email                 string    `json:"email"`                    // email
	CreatedOn             time.Time `json:"created_on"`               // created_on
	ArchivedOn            *NullTime `json:"archived_on"`              // archived_on
	FirstName             string    `json:"first_name"`               // first_name
	UpdatedOn             *NullTime `json:"updated_on"`               // updated_on
	ID                    uint64    `json:"id"`                       // id
	Username              string    `json:"username"`                 // username
	PasswordLastChangedOn *NullTime `json:"password_last_changed_on"` // password_last_changed_on
	Salt                  []byte    `json:"salt"`                     // salt
	LastName              string    `json:"last_name"`                // last_name
	IsAdmin               bool      `json:"is_admin"`                 // is_admin
	Password              string    `json:"password"`                 // password
}

type UserListResponse struct {
	ListResponse
	Users []User `json:"users"`
}
