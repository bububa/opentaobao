package traveltrade

import (
	"sync"
)

// ProductPriceStockDto 结构体
type ProductPriceStockDto struct {
	// 场次价库信息
	Sessions []ProductSessionDto `json:"sessions,omitempty" xml:"sessions>product_session_dto,omitempty"`
	// 日期。yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 产品结算单价。单位：分
	WholesalePrice int64 `json:"wholesale_price,omitempty" xml:"wholesale_price,omitempty"`
	// 产品销售单价。单位：分
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 是否可售卖；true：可售卖
	CanSell bool `json:"can_sell,omitempty" xml:"can_sell,omitempty"`
}

var poolProductPriceStockDto = sync.Pool{
	New: func() any {
		return new(ProductPriceStockDto)
	},
}

// GetProductPriceStockDto() 从对象池中获取ProductPriceStockDto
func GetProductPriceStockDto() *ProductPriceStockDto {
	return poolProductPriceStockDto.Get().(*ProductPriceStockDto)
}

// ReleaseProductPriceStockDto 释放ProductPriceStockDto
func ReleaseProductPriceStockDto(v *ProductPriceStockDto) {
	v.Sessions = v.Sessions[:0]
	v.Date = ""
	v.Stock = 0
	v.WholesalePrice = 0
	v.RetailPrice = 0
	v.CanSell = false
	poolProductPriceStockDto.Put(v)
}
