package wms

import (
	"sync"
)

// CainiaoInventoryProfitlossOrderitemlist 结构体
type CainiaoInventoryProfitlossOrderitemlist struct {
	// 损益详情
	OrderItem *CainiaoInventoryProfitlossOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolCainiaoInventoryProfitlossOrderitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoInventoryProfitlossOrderitemlist)
	},
}

// GetCainiaoInventoryProfitlossOrderitemlist() 从对象池中获取CainiaoInventoryProfitlossOrderitemlist
func GetCainiaoInventoryProfitlossOrderitemlist() *CainiaoInventoryProfitlossOrderitemlist {
	return poolCainiaoInventoryProfitlossOrderitemlist.Get().(*CainiaoInventoryProfitlossOrderitemlist)
}

// ReleaseCainiaoInventoryProfitlossOrderitemlist 释放CainiaoInventoryProfitlossOrderitemlist
func ReleaseCainiaoInventoryProfitlossOrderitemlist(v *CainiaoInventoryProfitlossOrderitemlist) {
	v.OrderItem = nil
	poolCainiaoInventoryProfitlossOrderitemlist.Put(v)
}
