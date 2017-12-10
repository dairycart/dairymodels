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
	ArchivedOn         NullTime  `json:"archived_on"`          // archived_on
	OptionSummary      string    `json:"option_summary"`       // option_summary
	UpdatedOn          NullTime  `json:"updated_on"`           // updated_on
	UPC                string    `json:"upc"`                  // upc

	// useful for responses
	ApplicableOptionValues []ProductOptionValue `json:"applicable_options,omitempty"`
}

// ProductCreationInput is a struct that represents a product creation body
type ProductCreationInput struct {
	ProductWidth       float64                      `json:"product_width"`        // product_width
	PackageLength      float64                      `json:"package_length"`       // package_length
	SalePrice          float64                      `json:"sale_price"`           // sale_price
	Description        string                       `json:"description"`          // description
	PackageWeight      float64                      `json:"package_weight"`       // package_weight
	Price              float64                      `json:"price"`                // price
	ProductWeight      float64                      `json:"product_weight"`       // product_weight
	Quantity           uint32                       `json:"quantity"`             // quantity
	ProductHeight      float64                      `json:"product_height"`       // product_height
	Taxable            bool                         `json:"taxable"`              // taxable
	Brand              string                       `json:"brand"`                // brand
	ProductLength      float64                      `json:"product_length"`       // product_length
	AvailableOn        time.Time                    `json:"available_on"`         // available_on
	QuantityPerPackage uint32                       `json:"quantity_per_package"` // quantity_per_package
	OnSale             bool                         `json:"on_sale"`              // on_sale
	Name               string                       `json:"name"`                 // name
	SKU                string                       `json:"sku"`                  // sku
	Manufacturer       string                       `json:"manufacturer"`         // manufacturer
	Subtitle           string                       `json:"subtitle"`             // subtitle
	PackageWidth       float64                      `json:"package_width"`        // package_width
	Cost               float64                      `json:"cost"`                 // cost
	PackageHeight      float64                      `json:"package_height"`       // package_height
	OptionSummary      string                       `json:"option_summary"`       // option_summary
	UpdatedOn          NullTime                     `json:"updated_on"`           // updated_on
	UPC                string                       `json:"upc"`                  // upc
	Options            []ProductOptionCreationInput `json:"options"`
}
