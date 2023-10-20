package wms

import (
	"sync"
)

// CainiaoStockInBillOrderitem 结构体
type CainiaoStockInBillOrderitem struct {
	// 仓库收货商品信息
	InventoryItemList []CainiaoStockInBillInventoryitemlist `json:"inventory_item_list,omitempty" xml:"inventory_item_list>cainiao_stock_in_bill_inventoryitemlist,omitempty"`
	// ERP订单明细ID
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 通知数量
	ApplyQty int64 `json:"apply_qty,omitempty" xml:"apply_qty,omitempty"`
}

var poolCainiaoStockInBillOrderitem = sync.Pool{
	New: func() any {
		return new(CainiaoStockInBillOrderitem)
	},
}

// GetCainiaoStockInBillOrderitem() 从对象池中获取CainiaoStockInBillOrderitem
func GetCainiaoStockInBillOrderitem() *CainiaoStockInBillOrderitem {
	return poolCainiaoStockInBillOrderitem.Get().(*CainiaoStockInBillOrderitem)
}

// ReleaseCainiaoStockInBillOrderitem 释放CainiaoStockInBillOrderitem
func ReleaseCainiaoStockInBillOrderitem(v *CainiaoStockInBillOrderitem) {
	v.InventoryItemList = v.InventoryItemList[:0]
	v.OrderItemId = ""
	v.ItemId = ""
	v.ItemCode = ""
	v.ApplyQty = 0
	poolCainiaoStockInBillOrderitem.Put(v)
}
