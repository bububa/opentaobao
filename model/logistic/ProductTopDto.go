package logistic

import (
	"sync"
)

// ProductTopDto 结构体
type ProductTopDto struct {
	// price
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// Actual dimension of a sku in its original packaging. Default unit: centimeters
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// Actual dimension of a sku in its original packaging. Default unit: centimeters
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// Total weight of a SKU in its original packaging. Default unit: kilograms
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// Actual dimension of a sku in its original packaging. Default unit: centimeters
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// AE sku_id to identify a unit of sku
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// Quantity of a sku in the order. It&#39;s used to calculate the total number of products in a parcel
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolProductTopDto = sync.Pool{
	New: func() any {
		return new(ProductTopDto)
	},
}

// GetProductTopDto() 从对象池中获取ProductTopDto
func GetProductTopDto() *ProductTopDto {
	return poolProductTopDto.Get().(*ProductTopDto)
}

// ReleaseProductTopDto 释放ProductTopDto
func ReleaseProductTopDto(v *ProductTopDto) {
	v.Price = ""
	v.Length = ""
	v.Width = ""
	v.Weight = ""
	v.Height = ""
	v.SkuId = ""
	v.Quantity = 0
	poolProductTopDto.Put(v)
}
