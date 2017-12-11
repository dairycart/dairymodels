package models

// LoginAttempt represents a Dairycart login attempt
type LoginAttempt struct {
	ID         uint64     `json:"id"`         // id
	Username   string     `json:"username"`   // username
	Successful bool       `json:"successful"` // successful
	CreatedOn  *Dairytime `json:"created_on"` // created_on
}

type LoginAttemptCreationInput struct {
	ID         uint64 `json:"id,omitempty"`         // id
	Username   string `json:"username,omitempty"`   // username
	Successful bool   `json:"successful,omitempty"` // successful

}

type LoginAttemptListResponse struct {
	ListResponse
	LoginAttempts []LoginAttempt `json:"login_attempts"`
}
