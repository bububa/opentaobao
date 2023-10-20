package happytrip

import (
	"sync"
)

// OrderGetResult 结构体
type OrderGetResult struct {
	// 费用信息，如果订单有折扣，这里为折后价格，如果没有折扣，这里和original_price保持一致
	Price *CostInfo `json:"price,omitempty" xml:"price,omitempty"`
	// 订单信息
	Order *OrderInfo `json:"order,omitempty" xml:"order,omitempty"`
	// 拼车信息
	Carpool *CarpoolInfo `json:"carpool,omitempty" xml:"carpool,omitempty"`
	// 订单改派信息
	ReassignInfo *ReassignInfo `json:"reassign_info,omitempty" xml:"reassign_info,omitempty"`
}

var poolOrderGetResult = sync.Pool{
	New: func() any {
		return new(OrderGetResult)
	},
}

// GetOrderGetResult() 从对象池中获取OrderGetResult
func GetOrderGetResult() *OrderGetResult {
	return poolOrderGetResult.Get().(*OrderGetResult)
}

// ReleaseOrderGetResult 释放OrderGetResult
func ReleaseOrderGetResult(v *OrderGetResult) {
	v.Price = nil
	v.Order = nil
	v.Carpool = nil
	v.ReassignInfo = nil
	poolOrderGetResult.Put(v)
}
