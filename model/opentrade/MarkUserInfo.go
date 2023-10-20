package opentrade

import (
	"sync"
)

// MarkUserInfo 结构体
type MarkUserInfo struct {
	// 用户openId
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	// 用户状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品SKU ID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 专属下单商品数量
	Quality int64 `json:"quality,omitempty" xml:"quality,omitempty"`
}

var poolMarkUserInfo = sync.Pool{
	New: func() any {
		return new(MarkUserInfo)
	},
}

// GetMarkUserInfo() 从对象池中获取MarkUserInfo
func GetMarkUserInfo() *MarkUserInfo {
	return poolMarkUserInfo.Get().(*MarkUserInfo)
}

// ReleaseMarkUserInfo 释放MarkUserInfo
func ReleaseMarkUserInfo(v *MarkUserInfo) {
	v.UserOpenId = ""
	v.Status = ""
	v.ItemId = 0
	v.SkuId = 0
	v.Quality = 0
	poolMarkUserInfo.Put(v)
}
