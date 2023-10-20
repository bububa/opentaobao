package tmallhk

import (
	"sync"
)

// SupplierBreakOrderRequest 结构体
type SupplierBreakOrderRequest struct {
	// 毁单商品详细信息
	BrokenOrderItemInfos []BrokenOrderItemInfo `json:"broken_order_item_infos,omitempty" xml:"broken_order_item_infos>broken_order_item_info,omitempty"`
	// 订单毁单时间
	BreakOrderTime string `json:"break_order_time,omitempty" xml:"break_order_time,omitempty"`
	// 毁单操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 主订单信息
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolSupplierBreakOrderRequest = sync.Pool{
	New: func() any {
		return new(SupplierBreakOrderRequest)
	},
}

// GetSupplierBreakOrderRequest() 从对象池中获取SupplierBreakOrderRequest
func GetSupplierBreakOrderRequest() *SupplierBreakOrderRequest {
	return poolSupplierBreakOrderRequest.Get().(*SupplierBreakOrderRequest)
}

// ReleaseSupplierBreakOrderRequest 释放SupplierBreakOrderRequest
func ReleaseSupplierBreakOrderRequest(v *SupplierBreakOrderRequest) {
	v.BrokenOrderItemInfos = v.BrokenOrderItemInfos[:0]
	v.BreakOrderTime = ""
	v.Operator = ""
	v.OrderId = 0
	poolSupplierBreakOrderRequest.Put(v)
}
