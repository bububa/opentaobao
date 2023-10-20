package axindata

import (
	"sync"
)

// FscRouteProjectPriceApiDto 结构体
type FscRouteProjectPriceApiDto struct {
	// 价格类型 ADULT:成人价 CHILD:儿童价
	PriceCategory string `json:"price_category,omitempty" xml:"price_category,omitempty"`
	// 同行价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 加返
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 门市价
	SalePrice int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
}

var poolFscRouteProjectPriceApiDto = sync.Pool{
	New: func() any {
		return new(FscRouteProjectPriceApiDto)
	},
}

// GetFscRouteProjectPriceApiDto() 从对象池中获取FscRouteProjectPriceApiDto
func GetFscRouteProjectPriceApiDto() *FscRouteProjectPriceApiDto {
	return poolFscRouteProjectPriceApiDto.Get().(*FscRouteProjectPriceApiDto)
}

// ReleaseFscRouteProjectPriceApiDto 释放FscRouteProjectPriceApiDto
func ReleaseFscRouteProjectPriceApiDto(v *FscRouteProjectPriceApiDto) {
	v.PriceCategory = ""
	v.Price = 0
	v.PromotionPrice = 0
	v.SalePrice = 0
	poolFscRouteProjectPriceApiDto.Put(v)
}
