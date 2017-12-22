package models

import (
	"time"
)

// ProductImage represents a Dairycart product image
type ProductImage struct {
	ID         uint64     `json:"id"`          // id
	ProductID  uint64     `json:"product_id"`  // product_id
	ImageURL   string     `json:"image_url"`   // image_url
	CreatedOn  time.Time  `json:"created_on"`  // created_on
	UpdatedOn  *Dairytime `json:"updated_on"`  // updated_on
	ArchivedOn *Dairytime `json:"archived_on"` // archived_on
}

// ProductImageCreationInput is a struct to use for creating ProductImages
type ProductImageCreationInput struct {
	ProductID uint64 `json:"product_id,omitempty"` // product_id
	ImageURL  string `json:"image_url,omitempty"`  // image_url
}

// ProductImageUpdateInput is a struct to use for updating ProductImages
type ProductImageUpdateInput struct {
	ProductID uint64 `json:"product_id,omitempty"` // product_id
	ImageURL  string `json:"image_url,omitempty"`  // image_url
}

type ProductImageListResponse struct {
	ListResponse
	ProductImages []ProductImage `json:"product_images"`
}
