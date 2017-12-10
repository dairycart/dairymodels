package models

import (
	"time"
)

// ProductRoot represents a Dairycart product root
type ProductRoot struct {
	AvailableOn        time.Time `json:"available_on"`         // available_on
	ProductLength      float64   `json:"product_length"`       // product_length
	UpdatedOn          NullTime  `json:"updated_on"`           // updated_on
	SKUPrefix          string    `json:"sku_prefix"`           // sku_prefix
	PackageHeight      float64   `json:"package_height"`       // package_height
	ProductWeight      float64   `json:"product_weight"`       // product_weight
	ProductWidth       float64   `json:"product_width"`        // product_width
	QuantityPerPackage uint32    `json:"quantity_per_package"` // quantity_per_package
	Name               string    `json:"name"`                 // name
	ProductHeight      float64   `json:"product_height"`       // product_height
	PackageLength      float64   `json:"package_length"`       // package_length
	CreatedOn          time.Time `json:"created_on"`           // created_on
	Cost               float64   `json:"cost"`                 // cost
	Brand              string    `json:"brand"`                // brand
	Subtitle           string    `json:"subtitle"`             // subtitle
	PackageWeight      float64   `json:"package_weight"`       // package_weight
	ArchivedOn         NullTime  `json:"archived_on"`          // archived_on
	ID                 uint64    `json:"id"`                   // id
	PackageWidth       float64   `json:"package_width"`        // package_width
	Description        string    `json:"description"`          // description
	Manufacturer       string    `json:"manufacturer"`         // manufacturer
	Taxable            bool      `json:"taxable"`              // taxable

	// useful for responses
	Options  []ProductOption `json:"options"`
	Products []Product       `json:"products"`
}
