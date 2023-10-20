package aedropshiper

import (
	"sync"
)

// PlaceOrderRequest4OpenApiDto 结构体
type PlaceOrderRequest4OpenApiDto struct {
	// 商品属性
	ProductItems []ProductBaseItem `json:"product_items,omitempty" xml:"product_items>product_base_item,omitempty"`
	// 物流地址信息
	LogisticsAddress *MaillingAddressRequestDto `json:"logistics_address,omitempty" xml:"logistics_address,omitempty"`
}

var poolPlaceOrderRequest4OpenApiDto = sync.Pool{
	New: func() any {
		return new(PlaceOrderRequest4OpenApiDto)
	},
}

// GetPlaceOrderRequest4OpenApiDto() 从对象池中获取PlaceOrderRequest4OpenApiDto
func GetPlaceOrderRequest4OpenApiDto() *PlaceOrderRequest4OpenApiDto {
	return poolPlaceOrderRequest4OpenApiDto.Get().(*PlaceOrderRequest4OpenApiDto)
}

// ReleasePlaceOrderRequest4OpenApiDto 释放PlaceOrderRequest4OpenApiDto
func ReleasePlaceOrderRequest4OpenApiDto(v *PlaceOrderRequest4OpenApiDto) {
	v.ProductItems = v.ProductItems[:0]
	v.LogisticsAddress = nil
	poolPlaceOrderRequest4OpenApiDto.Put(v)
}
