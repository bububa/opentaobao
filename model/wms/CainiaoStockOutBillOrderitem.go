package wms

import (
	"sync"
)

// CainiaoStockOutBillOrderitem 结构体
type CainiaoStockOutBillOrderitem struct {
	// 商品属性列表
	InventoryItemList []CainiaoStockOutBillInventoryitemlist `json:"inventory_item_list,omitempty" xml:"inventory_item_list>cainiao_stock_out_bill_inventoryitemlist,omitempty"`
	// order_item_id
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 通知数量
	ApplyQty int64 `json:"apply_qty,omitempty" xml:"apply_qty,omitempty"`
}

var poolCainiaoStockOutBillOrderitem = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillOrderitem)
	},
}

// GetCainiaoStockOutBillOrderitem() 从对象池中获取CainiaoStockOutBillOrderitem
func GetCainiaoStockOutBillOrderitem() *CainiaoStockOutBillOrderitem {
	return poolCainiaoStockOutBillOrderitem.Get().(*CainiaoStockOutBillOrderitem)
}

// ReleaseCainiaoStockOutBillOrderitem 释放CainiaoStockOutBillOrderitem
func ReleaseCainiaoStockOutBillOrderitem(v *CainiaoStockOutBillOrderitem) {
	v.InventoryItemList = v.InventoryItemList[:0]
	v.OrderItemId = ""
	v.ItemId = ""
	v.ItemCode = ""
	v.ApplyQty = 0
	poolCainiaoStockOutBillOrderitem.Put(v)
}
