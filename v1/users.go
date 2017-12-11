package models

// User represents a Dairycart user
type User struct {
	ID                    uint64     `json:"id"`                       // id
	FirstName             string     `json:"first_name"`               // first_name
	LastName              string     `json:"last_name"`                // last_name
	Username              string     `json:"username"`                 // username
	Email                 string     `json:"email"`                    // email
	Password              string     `json:"password"`                 // password
	Salt                  []byte     `json:"salt"`                     // salt
	IsAdmin               bool       `json:"is_admin"`                 // is_admin
	PasswordLastChangedOn *Dairytime `json:"password_last_changed_on"` // password_last_changed_on
	CreatedOn             *Dairytime `json:"created_on"`               // created_on
	UpdatedOn             *Dairytime `json:"updated_on"`               // updated_on
	ArchivedOn            *Dairytime `json:"archived_on"`              // archived_on
}

type UserCreationInput struct {
	ID                    uint64     `json:"id,omitempty"`                       // id
	FirstName             string     `json:"first_name,omitempty"`               // first_name
	LastName              string     `json:"last_name,omitempty"`                // last_name
	Username              string     `json:"username,omitempty"`                 // username
	Email                 string     `json:"email,omitempty"`                    // email
	Password              string     `json:"password,omitempty"`                 // password
	Salt                  []byte     `json:"salt,omitempty"`                     // salt
	IsAdmin               bool       `json:"is_admin,omitempty"`                 // is_admin
	PasswordLastChangedOn *Dairytime `json:"password_last_changed_on,omitempty"` // password_last_changed_on

}

type UserListResponse struct {
	ListResponse
	Users []User `json:"users"`
}
