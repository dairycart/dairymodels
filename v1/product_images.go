package models

import (
	"time"
)

// ProductImage represents a Dairycart product image
type ProductImage struct {
	ID           uint64     `json:"id"`            // id
	ProductID    uint64     `json:"product_id"`    // product_id
	ThumbnailURL string     `json:"thumbnail_url"` // thumbnail_url
	MainURL      string     `json:"main_url"`      // main_url
	OriginalURL  string     `json:"original_url"`  // original_url
	SourceURL    string     `json:"source_url"`    // source_url
	CreatedOn    time.Time  `json:"created_on"`    // created_on
	UpdatedOn    *Dairytime `json:"updated_on"`    // updated_on
	ArchivedOn   *Dairytime `json:"archived_on"`   // archived_on
}

// ProductImageCreationInput is a struct to use for creating ProductImages
type ProductImageCreationInput struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// ProductImageUpdateInput is a struct to use for updating ProductImages
type ProductImageUpdateInput struct {
	ProductID    uint64 `json:"product_id,omitempty"`    // product_id
	ThumbnailURL string `json:"thumbnail_url,omitempty"` // thumbnail_url
	MainURL      string `json:"main_url,omitempty"`      // main_url
	OriginalURL  string `json:"original_url,omitempty"`  // original_url
}

type ProductImageListResponse struct {
	ListResponse
	ProductImages []ProductImage `json:"product_images"`
}
