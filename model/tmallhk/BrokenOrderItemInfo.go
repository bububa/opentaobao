package tmallhk

import (
	"sync"
)

// BrokenOrderItemInfo 结构体
type BrokenOrderItemInfo struct {
	// 毁单商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 毁单商品下架时间
	ItemBrokenTime string `json:"item_broken_time,omitempty" xml:"item_broken_time,omitempty"`
	// 毁单商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolBrokenOrderItemInfo = sync.Pool{
	New: func() any {
		return new(BrokenOrderItemInfo)
	},
}

// GetBrokenOrderItemInfo() 从对象池中获取BrokenOrderItemInfo
func GetBrokenOrderItemInfo() *BrokenOrderItemInfo {
	return poolBrokenOrderItemInfo.Get().(*BrokenOrderItemInfo)
}

// ReleaseBrokenOrderItemInfo 释放BrokenOrderItemInfo
func ReleaseBrokenOrderItemInfo(v *BrokenOrderItemInfo) {
	v.ItemName = ""
	v.ItemBrokenTime = ""
	v.ItemId = 0
	poolBrokenOrderItemInfo.Put(v)
}
