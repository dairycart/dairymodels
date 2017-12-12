package models

// User represents a Dairycart user
type User struct {
	Email                 string     `json:"email"`                    // email
	CreatedOn             *Dairytime `json:"created_on"`               // created_on
	ArchivedOn            *Dairytime `json:"archived_on"`              // archived_on
	FirstName             string     `json:"first_name"`               // first_name
	UpdatedOn             *Dairytime `json:"updated_on"`               // updated_on
	ID                    uint64     `json:"id"`                       // id
	Username              string     `json:"username"`                 // username
	PasswordLastChangedOn *Dairytime `json:"password_last_changed_on"` // password_last_changed_on
	Salt                  []byte     `json:"salt"`                     // salt
	LastName              string     `json:"last_name"`                // last_name
	IsAdmin               bool       `json:"is_admin"`                 // is_admin
	Password              string     `json:"password"`                 // password
}

// UserCreationInput is a struct to use for creating product options
type UserCreationInput struct {
	Email     string `json:"email,omitempty"`      // email
	FirstName string `json:"first_name,omitempty"` // first_name
	Username  string `json:"username,omitempty"`   // username
	LastName  string `json:"last_name,omitempty"`  // last_name
	IsAdmin   bool   `json:"is_admin,omitempty"`   // is_admin
	Password  string `json:"password,omitempty"`   // password
}

// UserUpdateInput is a struct to use for updating Users
type UserUpdateInput struct {
	Email           string `json:"email,omitempty"`            // email
	FirstName       string `json:"first_name,omitempty"`       // first_name
	Username        string `json:"username,omitempty"`         // username
	LastName        string `json:"last_name,omitempty"`        // last_name
	IsAdmin         bool   `json:"is_admin,omitempty"`         // is_admin
	NewPassword     string `json:"new_password,omitempty"`     // password
	CurrentPassword string `json:"current_password,omitempty"` // password
}

type UserListResponse struct {
	ListResponse
	Users []User `json:"users"`
}
