package aedropshiper

// PackageInfoDto 结构体
type PackageInfoDto struct {
	// The gross weight of the product
	GrossWeight string `json:"gross_weight,omitempty" xml:"gross_weight,omitempty"`
	// The length of the product
	PackageLength int64 `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// Product height
	PackageHeight int64 `json:"package_height,omitempty" xml:"package_height,omitempty"`
	// Product width
	PackageWidth int64 `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// Number of basic products for custom weighing
	BaseUnit int64 `json:"base_unit,omitempty" xml:"base_unit,omitempty"`
	// Unit of commodity
	ProductUnit int64 `json:"product_unit,omitempty" xml:"product_unit,omitempty"`
	// Type of packaging
	PackageType bool `json:"package_type,omitempty" xml:"package_type,omitempty"`
}
