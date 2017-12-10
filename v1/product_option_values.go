package models

import (
	"time"
)

// ProductOptionValue represents a Dairycart product option value
type ProductOptionValue struct {
	Value           string    `json:"value"`             // value
	CreatedOn       time.Time `json:"created_on"`        // created_on
	ID              uint64    `json:"id"`                // id
	ArchivedOn      NullTime  `json:"archived_on"`       // archived_on
	UpdatedOn       NullTime  `json:"updated_on"`        // updated_on
	ProductOptionID uint64    `json:"product_option_id"` // product_option_id
}
