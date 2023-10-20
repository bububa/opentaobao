package wms

import (
	"sync"
)

// CainiaoReturnBillInventoryitemlist 结构体
type CainiaoReturnBillInventoryitemlist struct {
	// 订单详情
	InventoryItem *CainiaoReturnBillInventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}

var poolCainiaoReturnBillInventoryitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoReturnBillInventoryitemlist)
	},
}

// GetCainiaoReturnBillInventoryitemlist() 从对象池中获取CainiaoReturnBillInventoryitemlist
func GetCainiaoReturnBillInventoryitemlist() *CainiaoReturnBillInventoryitemlist {
	return poolCainiaoReturnBillInventoryitemlist.Get().(*CainiaoReturnBillInventoryitemlist)
}

// ReleaseCainiaoReturnBillInventoryitemlist 释放CainiaoReturnBillInventoryitemlist
func ReleaseCainiaoReturnBillInventoryitemlist(v *CainiaoReturnBillInventoryitemlist) {
	v.InventoryItem = nil
	poolCainiaoReturnBillInventoryitemlist.Put(v)
}
