package wms

import (
	"sync"
)

// CainiaoStockOutBillPackageitem 结构体
type CainiaoStockOutBillPackageitem struct {
	// ERP订单明细ID
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 菜鸟商品编码
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// ERP商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 库存类型1 可销售库存 101残次品
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 数量
	ItemQty int64 `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
}

var poolCainiaoStockOutBillPackageitem = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillPackageitem)
	},
}

// GetCainiaoStockOutBillPackageitem() 从对象池中获取CainiaoStockOutBillPackageitem
func GetCainiaoStockOutBillPackageitem() *CainiaoStockOutBillPackageitem {
	return poolCainiaoStockOutBillPackageitem.Get().(*CainiaoStockOutBillPackageitem)
}

// ReleaseCainiaoStockOutBillPackageitem 释放CainiaoStockOutBillPackageitem
func ReleaseCainiaoStockOutBillPackageitem(v *CainiaoStockOutBillPackageitem) {
	v.OrderItemId = ""
	v.ItemId = ""
	v.ItemCode = ""
	v.InventoryType = 0
	v.ItemQty = 0
	poolCainiaoStockOutBillPackageitem.Put(v)
}
