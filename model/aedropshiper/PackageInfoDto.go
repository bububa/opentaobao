package aedropshiper

import (
	"sync"
)

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

var poolPackageInfoDto = sync.Pool{
	New: func() any {
		return new(PackageInfoDto)
	},
}

// GetPackageInfoDto() 从对象池中获取PackageInfoDto
func GetPackageInfoDto() *PackageInfoDto {
	return poolPackageInfoDto.Get().(*PackageInfoDto)
}

// ReleasePackageInfoDto 释放PackageInfoDto
func ReleasePackageInfoDto(v *PackageInfoDto) {
	v.GrossWeight = ""
	v.PackageLength = 0
	v.PackageHeight = 0
	v.PackageWidth = 0
	v.BaseUnit = 0
	v.ProductUnit = 0
	v.PackageType = false
	poolPackageInfoDto.Put(v)
}
