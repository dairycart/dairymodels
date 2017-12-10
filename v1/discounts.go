package models

import (
	"time"
)

// Discount represents a Dairycart discount
type Discount struct {
	ID            uint64    `json:"id"`             // id
	Amount        float64   `json:"amount"`         // amount
	LimitedUse    bool      `json:"limited_use"`    // limited_use
	RequiresCode  bool      `json:"requires_code"`  // requires_code
	DiscountType  string    `json:"discount_type"`  // discount_type
	ArchivedOn    *NullTime `json:"archived_on"`    // archived_on
	Code          string    `json:"code"`           // code
	ExpiresOn     *NullTime `json:"expires_on"`     // expires_on
	UpdatedOn     *NullTime `json:"updated_on"`     // updated_on
	StartsOn      time.Time `json:"starts_on"`      // starts_on
	CreatedOn     time.Time `json:"created_on"`     // created_on
	NumberOfUses  uint64    `json:"number_of_uses"` // number_of_uses
	Name          string    `json:"name"`           // name
	LoginRequired bool      `json:"login_required"` // login_required
}

type DiscountListResponse struct {
	ListResponse
	Discounts []Discount `json:"discounts"`
}
