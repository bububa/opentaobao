package wms

import (
	"sync"
)

// CainiaoStockOutBillOrderitemlist 结构体
type CainiaoStockOutBillOrderitemlist struct {
	// 订单商品信息
	OrderItem *CainiaoStockOutBillOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolCainiaoStockOutBillOrderitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillOrderitemlist)
	},
}

// GetCainiaoStockOutBillOrderitemlist() 从对象池中获取CainiaoStockOutBillOrderitemlist
func GetCainiaoStockOutBillOrderitemlist() *CainiaoStockOutBillOrderitemlist {
	return poolCainiaoStockOutBillOrderitemlist.Get().(*CainiaoStockOutBillOrderitemlist)
}

// ReleaseCainiaoStockOutBillOrderitemlist 释放CainiaoStockOutBillOrderitemlist
func ReleaseCainiaoStockOutBillOrderitemlist(v *CainiaoStockOutBillOrderitemlist) {
	v.OrderItem = nil
	poolCainiaoStockOutBillOrderitemlist.Put(v)
}
