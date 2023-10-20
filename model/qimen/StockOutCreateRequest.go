package qimen

import (
	"sync"
)

// StockOutCreateRequest 结构体
type StockOutCreateRequest struct {
	// 单据信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 出库单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockoutCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolStockOutCreateRequest = sync.Pool{
	New: func() any {
		return new(StockOutCreateRequest)
	},
}

// GetStockOutCreateRequest() 从对象池中获取StockOutCreateRequest
func GetStockOutCreateRequest() *StockOutCreateRequest {
	return poolStockOutCreateRequest.Get().(*StockOutCreateRequest)
}

// ReleaseStockOutCreateRequest 释放StockOutCreateRequest
func ReleaseStockOutCreateRequest(v *StockOutCreateRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.DeliveryOrder = nil
	v.ExtendProps = nil
	poolStockOutCreateRequest.Put(v)
}
