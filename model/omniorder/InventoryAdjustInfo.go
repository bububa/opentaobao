package omniorder

import (
	"sync"
)

// InventoryAdjustInfo 结构体
type InventoryAdjustInfo struct {
	// 需要调整的原始门店ID
	OriginalWarehouseId string `json:"original_warehouse_id,omitempty" xml:"original_warehouse_id,omitempty"`
	// 淘宝订单号
	TbTradeOrder string `json:"tb_trade_order,omitempty" xml:"tb_trade_order,omitempty"`
	// 流水号
	BillNum string `json:"bill_num,omitempty" xml:"bill_num,omitempty"`
	// 库存类型
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 调整类型
	AdjustType string `json:"adjust_type,omitempty" xml:"adjust_type,omitempty"`
	// 淘宝子订单号
	TbSubTradeOrder string `json:"tb_sub_trade_order,omitempty" xml:"tb_sub_trade_order,omitempty"`
	// 需要调整到的目标门店ID
	TargetWarehouseId string `json:"target_warehouse_id,omitempty" xml:"target_warehouse_id,omitempty"`
	// ISV系统中商品编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 需调整的库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 淘宝后端商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 淘宝前端商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品的SKU编码
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolInventoryAdjustInfo = sync.Pool{
	New: func() any {
		return new(InventoryAdjustInfo)
	},
}

// GetInventoryAdjustInfo() 从对象池中获取InventoryAdjustInfo
func GetInventoryAdjustInfo() *InventoryAdjustInfo {
	return poolInventoryAdjustInfo.Get().(*InventoryAdjustInfo)
}

// ReleaseInventoryAdjustInfo 释放InventoryAdjustInfo
func ReleaseInventoryAdjustInfo(v *InventoryAdjustInfo) {
	v.OriginalWarehouseId = ""
	v.TbTradeOrder = ""
	v.BillNum = ""
	v.InventoryType = ""
	v.AdjustType = ""
	v.TbSubTradeOrder = ""
	v.TargetWarehouseId = ""
	v.OuterId = ""
	v.Quantity = 0
	v.ScItemId = 0
	v.ItemId = 0
	v.SkuId = 0
	poolInventoryAdjustInfo.Put(v)
}
