package ascp

import (
	"sync"
)

// HiErpItemInfo 结构体
type HiErpItemInfo struct {
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品唯一ID（预留给商家使用）
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 货品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolHiErpItemInfo = sync.Pool{
	New: func() any {
		return new(HiErpItemInfo)
	},
}

// GetHiErpItemInfo() 从对象池中获取HiErpItemInfo
func GetHiErpItemInfo() *HiErpItemInfo {
	return poolHiErpItemInfo.Get().(*HiErpItemInfo)
}

// ReleaseHiErpItemInfo 释放HiErpItemInfo
func ReleaseHiErpItemInfo(v *HiErpItemInfo) {
	v.ItemCode = ""
	v.Uid = ""
	v.ItemId = 0
	v.Quantity = 0
	poolHiErpItemInfo.Put(v)
}
