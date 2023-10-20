package wms

import (
	"sync"
)

// CainiaoReturnBillOrderitem 结构体
type CainiaoReturnBillOrderitem struct {
	// 商品信息
	InventoryItemList []CainiaoReturnBillInventoryitemlist `json:"inventory_item_list,omitempty" xml:"inventory_item_list>cainiao_return_bill_inventoryitemlist,omitempty"`
	// 商品ID
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}

var poolCainiaoReturnBillOrderitem = sync.Pool{
	New: func() any {
		return new(CainiaoReturnBillOrderitem)
	},
}

// GetCainiaoReturnBillOrderitem() 从对象池中获取CainiaoReturnBillOrderitem
func GetCainiaoReturnBillOrderitem() *CainiaoReturnBillOrderitem {
	return poolCainiaoReturnBillOrderitem.Get().(*CainiaoReturnBillOrderitem)
}

// ReleaseCainiaoReturnBillOrderitem 释放CainiaoReturnBillOrderitem
func ReleaseCainiaoReturnBillOrderitem(v *CainiaoReturnBillOrderitem) {
	v.InventoryItemList = v.InventoryItemList[:0]
	v.OrderItemId = ""
	v.ItemId = ""
	v.ItemCode = ""
	poolCainiaoReturnBillOrderitem.Put(v)
}
