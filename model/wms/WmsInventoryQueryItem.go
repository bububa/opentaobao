package wms

import (
	"sync"
)

// WmsInventoryQueryItem 结构体
type WmsInventoryQueryItem struct {
	// 失效日期，type=2时字段有返回值。
	DueDate string `json:"due_date,omitempty" xml:"due_date,omitempty"`
	// 生产日期，type=2时字段有返回值
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 库存批次号，type=2时字段有返回值。
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 渠道编码，type=3时字段有返回值。（TB 淘系，OTHERS 其他）
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 锁库存数量
	LockQuantity int64 `json:"lock_quantity,omitempty" xml:"lock_quantity,omitempty"`
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 库存类型(1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
}

var poolWmsInventoryQueryItem = sync.Pool{
	New: func() any {
		return new(WmsInventoryQueryItem)
	},
}

// GetWmsInventoryQueryItem() 从对象池中获取WmsInventoryQueryItem
func GetWmsInventoryQueryItem() *WmsInventoryQueryItem {
	return poolWmsInventoryQueryItem.Get().(*WmsInventoryQueryItem)
}

// ReleaseWmsInventoryQueryItem 释放WmsInventoryQueryItem
func ReleaseWmsInventoryQueryItem(v *WmsInventoryQueryItem) {
	v.DueDate = ""
	v.ProduceDate = ""
	v.BatchCode = ""
	v.ChannelCode = ""
	v.ItemId = ""
	v.StoreCode = ""
	v.LockQuantity = 0
	v.Quantity = 0
	v.InventoryType = 0
	poolWmsInventoryQueryItem.Put(v)
}
