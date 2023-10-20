package fenxiao

import (
	"sync"
)

// InventoryQueryResult 结构体
type InventoryQueryResult struct {
	// 查询成功列表
	ItemInventorys []InventoryInfoDetailDto `json:"item_inventorys,omitempty" xml:"item_inventorys>inventory_info_detail_dto,omitempty"`
	// tipInfos
	TipInfos []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
}

var poolInventoryQueryResult = sync.Pool{
	New: func() any {
		return new(InventoryQueryResult)
	},
}

// GetInventoryQueryResult() 从对象池中获取InventoryQueryResult
func GetInventoryQueryResult() *InventoryQueryResult {
	return poolInventoryQueryResult.Get().(*InventoryQueryResult)
}

// ReleaseInventoryQueryResult 释放InventoryQueryResult
func ReleaseInventoryQueryResult(v *InventoryQueryResult) {
	v.ItemInventorys = v.ItemInventorys[:0]
	v.TipInfos = v.TipInfos[:0]
	poolInventoryQueryResult.Put(v)
}
