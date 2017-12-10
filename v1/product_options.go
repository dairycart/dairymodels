package models

import (
	"time"
)

// ProductOption represents a Dairycart product option
type ProductOption struct {
	ProductRootID uint64    `json:"product_root_id"` // product_root_id
	CreatedOn     time.Time `json:"created_on"`      // created_on
	ID            uint64    `json:"id"`              // id
	ArchivedOn    NullTime  `json:"archived_on"`     // archived_on
	UpdatedOn     NullTime  `json:"updated_on"`      // updated_on
	Name          string    `json:"name"`            // name

	// useful for responses
	Values []ProductOptionValue `json:"values"`
}

type ProductOptionListResponse struct {
	ListResponse
	ProductOptions []ProductOption `json: "product_options"`
}

// ProductOptionCreationInput is a struct to use for creating product options
type ProductOptionCreationInput struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
