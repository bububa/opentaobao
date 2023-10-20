package ascp

import (
	"sync"
)

// HiErpModifyInventoryResp 结构体
type HiErpModifyInventoryResp struct {
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 当前库存数量
	CurrentNumber int64 `json:"current_number,omitempty" xml:"current_number,omitempty"`
}

var poolHiErpModifyInventoryResp = sync.Pool{
	New: func() any {
		return new(HiErpModifyInventoryResp)
	},
}

// GetHiErpModifyInventoryResp() 从对象池中获取HiErpModifyInventoryResp
func GetHiErpModifyInventoryResp() *HiErpModifyInventoryResp {
	return poolHiErpModifyInventoryResp.Get().(*HiErpModifyInventoryResp)
}

// ReleaseHiErpModifyInventoryResp 释放HiErpModifyInventoryResp
func ReleaseHiErpModifyInventoryResp(v *HiErpModifyInventoryResp) {
	v.StoreCode = ""
	v.ItemCode = ""
	v.ItemId = 0
	v.CurrentNumber = 0
	poolHiErpModifyInventoryResp.Put(v)
}
