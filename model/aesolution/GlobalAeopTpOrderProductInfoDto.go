package aesolution

import (
	"sync"
)

// GlobalAeopTpOrderProductInfoDto 结构体
type GlobalAeopTpOrderProductInfoDto struct {
	// product SKU details
	Sku string `json:"sku,omitempty" xml:"sku,omitempty"`
	// product name
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// Leaf category Id of the product
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// product quantity
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// product unit price
	UnitPrice *GlobalMoneyStr `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolGlobalAeopTpOrderProductInfoDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpOrderProductInfoDto)
	},
}

// GetGlobalAeopTpOrderProductInfoDto() 从对象池中获取GlobalAeopTpOrderProductInfoDto
func GetGlobalAeopTpOrderProductInfoDto() *GlobalAeopTpOrderProductInfoDto {
	return poolGlobalAeopTpOrderProductInfoDto.Get().(*GlobalAeopTpOrderProductInfoDto)
}

// ReleaseGlobalAeopTpOrderProductInfoDto 释放GlobalAeopTpOrderProductInfoDto
func ReleaseGlobalAeopTpOrderProductInfoDto(v *GlobalAeopTpOrderProductInfoDto) {
	v.Sku = ""
	v.ProductName = ""
	v.CategoryId = ""
	v.Quantity = 0
	v.UnitPrice = nil
	v.ProductId = 0
	poolGlobalAeopTpOrderProductInfoDto.Put(v)
}
