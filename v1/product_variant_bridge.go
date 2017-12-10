package models

import (
	"time"
)

// ProductVariantBridge represents a Dairycart product variant bridge
type ProductVariantBridge struct {
	ArchivedOn           NullTime  `json:"archived_on"`             // archived_on
	ProductID            uint64    `json:"product_id"`              // product_id
	ProductOptionValueID uint64    `json:"product_option_value_id"` // product_option_value_id
	CreatedOn            time.Time `json:"created_on"`              // created_on
	ID                   uint64    `json:"id"`                      // id
}

type ProductVariantBridgeListResponse struct {
	ListResponse
	ProductVariantBridges []ProductVariantBridge `json: "product_variant_bridge"`
}
