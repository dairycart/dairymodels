package models

import (
	"time"
)

// Product represents a Dairycart product
type Product struct {
	ProductWidth       float64   `json:"product_width"`        // product_width
	PackageLength      float64   `json:"package_length"`       // package_length
	SalePrice          float64   `json:"sale_price"`           // sale_price
	Description        string    `json:"description"`          // description
	PackageWeight      float64   `json:"package_weight"`       // package_weight
	Price              float64   `json:"price"`                // price
	ProductWeight      float64   `json:"product_weight"`       // product_weight
	Quantity           uint32    `json:"quantity"`             // quantity
	ProductRootID      uint64    `json:"product_root_id"`      // product_root_id
	ProductHeight      float64   `json:"product_height"`       // product_height
	Taxable            bool      `json:"taxable"`              // taxable
	Brand              string    `json:"brand"`                // brand
	ProductLength      float64   `json:"product_length"`       // product_length
	CreatedOn          time.Time `json:"created_on"`           // created_on
	AvailableOn        time.Time `json:"available_on"`         // available_on
	QuantityPerPackage uint32    `json:"quantity_per_package"` // quantity_per_package
	OnSale             bool      `json:"on_sale"`              // on_sale
	Name               string    `json:"name"`                 // name
	SKU                string    `json:"sku"`                  // sku
	Manufacturer       string    `json:"manufacturer"`         // manufacturer
	Subtitle           string    `json:"subtitle"`             // subtitle
	PackageWidth       float64   `json:"package_width"`        // package_width
	Cost               float64   `json:"cost"`                 // cost
	ID                 uint64    `json:"id"`                   // id
	PackageHeight      float64   `json:"package_height"`       // package_height
	ArchivedOn         *NullTime `json:"archived_on"`          // archived_on
	OptionSummary      string    `json:"option_summary"`       // option_summary
	UpdatedOn          *NullTime `json:"updated_on"`           // updated_on
	UPC                string    `json:"upc"`                  // upc

	// useful for responses
	ApplicableOptionValues []ProductOptionValue `json:"applicable_options,omitempty"`
}

type ProductListResponse struct {
	ListResponse
	Products []Product `json:"products"`
}

// ProductCreationInput is a struct that represents a product creation body
type ProductCreationInput struct {
	ProductWidth       float64                      `json:"product_width,omitempty"`        // product_width
	PackageLength      float64                      `json:"package_length,omitempty"`       // package_length
	SalePrice          float64                      `json:"sale_price,omitempty"`           // sale_price
	Description        string                       `json:"description,omitempty"`          // description
	PackageWeight      float64                      `json:"package_weight,omitempty"`       // package_weight
	Price              float64                      `json:"price,omitempty"`                // price
	ProductWeight      float64                      `json:"product_weight,omitempty"`       // product_weight
	Quantity           uint32                       `json:"quantity,omitempty"`             // quantity
	ProductHeight      float64                      `json:"product_height,omitempty"`       // product_height
	Taxable            bool                         `json:"taxable,omitempty"`              // taxable
	Brand              string                       `json:"brand,omitempty"`                // brand
	ProductLength      float64                      `json:"product_length,omitempty"`       // product_length
	AvailableOn        time.Time                    `json:"available_on,omitempty"`         // available_on
	QuantityPerPackage uint32                       `json:"quantity_per_package,omitempty"` // quantity_per_package
	OnSale             bool                         `json:"on_sale,omitempty"`              // on_sale
	Name               string                       `json:"name,omitempty"`                 // name
	SKU                string                       `json:"sku,omitempty"`                  // sku
	Manufacturer       string                       `json:"manufacturer,omitempty"`         // manufacturer
	Subtitle           string                       `json:"subtitle,omitempty"`             // subtitle
	PackageWidth       float64                      `json:"package_width,omitempty"`        // package_width
	Cost               float64                      `json:"cost,omitempty"`                 // cost
	PackageHeight      float64                      `json:"package_height,omitempty"`       // package_height
	OptionSummary      string                       `json:"option_summary,omitempty"`       // option_summary
	UpdatedOn          *NullTime                    `json:"updated_on,omitempty"`           // updated_on
	UPC                string                       `json:"upc,omitempty"`                  // upc
	Options            []ProductOptionCreationInput `json:"options,omitempty"`
}
