package wlb

import (
	"sync"
)

// WlbItemInventory 结构体
type WlbItemInventory struct {
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// SELLALBE 可销售库存&lt;br/&gt;DEFECTIVE 残次&lt;br/&gt;JISHUN 机损&lt;br/&gt;XIANGSHUN 箱损&lt;br/&gt;FREEZE 冻结库存&lt;br/&gt;ONROAD 在途库存
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 锁定库存数量
	LockQuantity int64 `json:"lock_quantity,omitempty" xml:"lock_quantity,omitempty"`
}

var poolWlbItemInventory = sync.Pool{
	New: func() any {
		return new(WlbItemInventory)
	},
}

// GetWlbItemInventory() 从对象池中获取WlbItemInventory
func GetWlbItemInventory() *WlbItemInventory {
	return poolWlbItemInventory.Get().(*WlbItemInventory)
}

// ReleaseWlbItemInventory 释放WlbItemInventory
func ReleaseWlbItemInventory(v *WlbItemInventory) {
	v.StoreCode = ""
	v.Type = ""
	v.ItemId = 0
	v.Quantity = 0
	v.LockQuantity = 0
	poolWlbItemInventory.Put(v)
}
