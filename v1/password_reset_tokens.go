package models

// PasswordResetToken represents a Dairycart password reset token
type PasswordResetToken struct {
	ID              uint64     `json:"id"`                // id
	UserID          uint64     `json:"user_id"`           // user_id
	Token           string     `json:"token"`             // token
	CreatedOn       *Dairytime `json:"created_on"`        // created_on
	ExpiresOn       *Dairytime `json:"expires_on"`        // expires_on
	PasswordResetOn *Dairytime `json:"password_reset_on"` // password_reset_on
}

type PasswordResetTokenCreationInput struct {
	ID              uint64     `json:"id,omitempty"`                // id
	UserID          uint64     `json:"user_id,omitempty"`           // user_id
	Token           string     `json:"token,omitempty"`             // token
	ExpiresOn       *Dairytime `json:"expires_on,omitempty"`        // expires_on
	PasswordResetOn *Dairytime `json:"password_reset_on,omitempty"` // password_reset_on

}

type PasswordResetTokenListResponse struct {
	ListResponse
	PasswordResetTokens []PasswordResetToken `json:"password_reset_tokens"`
}
