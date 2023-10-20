package qimen

import (
	"sync"
)

// ItemMapping 结构体
type ItemMapping struct {
	// 奇门仓储字段,C123,string(50),
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ShopNick string `json:"shopNick,omitempty" xml:"shopNick,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ItemSource string `json:"itemSource,omitempty" xml:"itemSource,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ShopItemId string `json:"shopItemId,omitempty" xml:"shopItemId,omitempty"`
	// 奇门仓储字段,C123,string(50),
	SkuId string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

var poolItemMapping = sync.Pool{
	New: func() any {
		return new(ItemMapping)
	},
}

// GetItemMapping() 从对象池中获取ItemMapping
func GetItemMapping() *ItemMapping {
	return poolItemMapping.Get().(*ItemMapping)
}

// ReleaseItemMapping 释放ItemMapping
func ReleaseItemMapping(v *ItemMapping) {
	v.OwnerCode = ""
	v.ShopNick = ""
	v.ItemSource = ""
	v.ItemId = ""
	v.ShopItemId = ""
	v.SkuId = ""
	poolItemMapping.Put(v)
}
