package wms

import (
	"sync"
)

// CainiaoStockOutBillInventoryitemlist 结构体
type CainiaoStockOutBillInventoryitemlist struct {
	// 商品属性
	InventoryItem *CainiaoStockOutBillInventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}

var poolCainiaoStockOutBillInventoryitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillInventoryitemlist)
	},
}

// GetCainiaoStockOutBillInventoryitemlist() 从对象池中获取CainiaoStockOutBillInventoryitemlist
func GetCainiaoStockOutBillInventoryitemlist() *CainiaoStockOutBillInventoryitemlist {
	return poolCainiaoStockOutBillInventoryitemlist.Get().(*CainiaoStockOutBillInventoryitemlist)
}

// ReleaseCainiaoStockOutBillInventoryitemlist 释放CainiaoStockOutBillInventoryitemlist
func ReleaseCainiaoStockOutBillInventoryitemlist(v *CainiaoStockOutBillInventoryitemlist) {
	v.InventoryItem = nil
	poolCainiaoStockOutBillInventoryitemlist.Put(v)
}
