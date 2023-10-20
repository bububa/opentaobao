package wms

import (
	"sync"
)

// CainiaoReturnBillOrderitemlist 结构体
type CainiaoReturnBillOrderitemlist struct {
	// 订单商品信息
	OrderItem *CainiaoReturnBillOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolCainiaoReturnBillOrderitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoReturnBillOrderitemlist)
	},
}

// GetCainiaoReturnBillOrderitemlist() 从对象池中获取CainiaoReturnBillOrderitemlist
func GetCainiaoReturnBillOrderitemlist() *CainiaoReturnBillOrderitemlist {
	return poolCainiaoReturnBillOrderitemlist.Get().(*CainiaoReturnBillOrderitemlist)
}

// ReleaseCainiaoReturnBillOrderitemlist 释放CainiaoReturnBillOrderitemlist
func ReleaseCainiaoReturnBillOrderitemlist(v *CainiaoReturnBillOrderitemlist) {
	v.OrderItem = nil
	poolCainiaoReturnBillOrderitemlist.Put(v)
}
