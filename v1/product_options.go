package models

// ProductOption represents a Dairycart product option
type ProductOption struct {
	ID            uint64     `json:"id"`              // id
	Name          string     `json:"name"`            // name
	ProductRootID uint64     `json:"product_root_id"` // product_root_id
	CreatedOn     *Dairytime `json:"created_on"`      // created_on
	UpdatedOn     *Dairytime `json:"updated_on"`      // updated_on
	ArchivedOn    *Dairytime `json:"archived_on"`     // archived_on

	// useful for responses
	Values []ProductOptionValue `json:"values"`
}

// ProductOptionCreationInput is a struct to use for creating product options
type ProductOptionCreationInput struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

type ProductOptionListResponse struct {
	ListResponse
	ProductOptions []ProductOption `json:"product_options"`
}
