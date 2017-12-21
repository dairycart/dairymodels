package models

import (
	"time"
)

// ProductVariantBridge represents a Dairycart product variant bridge
type ProductVariantBridge struct {
	ID                   uint64     `json:"id"`                      // id
	ProductID            uint64     `json:"product_id"`              // product_id
	ProductOptionValueID uint64     `json:"product_option_value_id"` // product_option_value_id
	CreatedOn            time.Time  `json:"created_on"`              // created_on
	ArchivedOn           *Dairytime `json:"archived_on"`             // archived_on
}

type ProductVariantBridgeCreationInput struct {
	ProductID            uint64 `json:"product_id,omitempty"`              // product_id
	ProductOptionValueID uint64 `json:"product_option_value_id,omitempty"` // product_option_value_id
}

// ProductVariantBridgeUpdateInput is a struct to use for updating ProductVariantBridges
type ProductVariantBridgeUpdateInput struct {
	ProductID            uint64 `json:"product_id,omitempty"`              // product_id
	ProductOptionValueID uint64 `json:"product_option_value_id,omitempty"` // product_option_value_id
}

type ProductVariantBridgeListResponse struct {
	ListResponse
	ProductVariantBridges []ProductVariantBridge `json:"product_variant_bridge"`
}
