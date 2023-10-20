package wms

import (
	"sync"
)

// CainiaoStockInBillInventoryitemlist 结构体
type CainiaoStockInBillInventoryitemlist struct {
	// 仓库收货商品信息
	InventoryItem *CainiaoStockInBillInventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}

var poolCainiaoStockInBillInventoryitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoStockInBillInventoryitemlist)
	},
}

// GetCainiaoStockInBillInventoryitemlist() 从对象池中获取CainiaoStockInBillInventoryitemlist
func GetCainiaoStockInBillInventoryitemlist() *CainiaoStockInBillInventoryitemlist {
	return poolCainiaoStockInBillInventoryitemlist.Get().(*CainiaoStockInBillInventoryitemlist)
}

// ReleaseCainiaoStockInBillInventoryitemlist 释放CainiaoStockInBillInventoryitemlist
func ReleaseCainiaoStockInBillInventoryitemlist(v *CainiaoStockInBillInventoryitemlist) {
	v.InventoryItem = nil
	poolCainiaoStockInBillInventoryitemlist.Put(v)
}
