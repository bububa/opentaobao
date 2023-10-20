package wms

import (
	"sync"
)

// CainiaoStockInBillOrderitemlist 结构体
type CainiaoStockInBillOrderitemlist struct {
	// 入库单信息
	OrderItem *CainiaoStockInBillOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolCainiaoStockInBillOrderitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoStockInBillOrderitemlist)
	},
}

// GetCainiaoStockInBillOrderitemlist() 从对象池中获取CainiaoStockInBillOrderitemlist
func GetCainiaoStockInBillOrderitemlist() *CainiaoStockInBillOrderitemlist {
	return poolCainiaoStockInBillOrderitemlist.Get().(*CainiaoStockInBillOrderitemlist)
}

// ReleaseCainiaoStockInBillOrderitemlist 释放CainiaoStockInBillOrderitemlist
func ReleaseCainiaoStockInBillOrderitemlist(v *CainiaoStockInBillOrderitemlist) {
	v.OrderItem = nil
	poolCainiaoStockInBillOrderitemlist.Put(v)
}
