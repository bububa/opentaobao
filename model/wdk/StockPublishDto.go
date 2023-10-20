package wdk

import (
	"sync"
)

// StockPublishDto 结构体
type StockPublishDto struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 更新数量，矢量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 订单号（商品粒度操作单）
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 订单类型，参见订单类型约束
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 订单描述
	OrderDesc string `json:"order_desc,omitempty" xml:"order_desc,omitempty"`
}

var poolStockPublishDto = sync.Pool{
	New: func() any {
		return new(StockPublishDto)
	},
}

// GetStockPublishDto() 从对象池中获取StockPublishDto
func GetStockPublishDto() *StockPublishDto {
	return poolStockPublishDto.Get().(*StockPublishDto)
}

// ReleaseStockPublishDto 释放StockPublishDto
func ReleaseStockPublishDto(v *StockPublishDto) {
	v.SkuCode = ""
	v.Quantity = ""
	v.OrderNo = ""
	v.OrderType = ""
	v.OrderDesc = ""
	poolStockPublishDto.Put(v)
}
