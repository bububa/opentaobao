package qimen

import (
	"sync"
)

// RequestDo 结构体
type RequestDo struct {
	// 奇门仓储字段
	WarehouseCodes []string `json:"warehouseCodes,omitempty" xml:"warehouseCodes>string,omitempty"`
	// 奇门仓储字段
	ItemCodes []string `json:"itemCodes,omitempty" xml:"itemCodes>string,omitempty"`
	// 奇门仓储字段
	ChannelCodes []string `json:"channelCodes,omitempty" xml:"channelCodes>string,omitempty"`
	// inventoryRules
	InventoryRules []InventoryRule `json:"inventoryRules,omitempty" xml:"inventoryRules>inventory_rule,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 奇门仓储字段,C123,string(50),必填,
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 奇门仓储字段,C123,string(50),必填,
	ShopNick string `json:"shopNick,omitempty" xml:"shopNick,omitempty"`
	// 奇门仓储字段,C123,string(50),必填,
	ItemSource string `json:"itemSource,omitempty" xml:"itemSource,omitempty"`
	// 奇门仓储字段,C123,string(50),必填,
	ShopItemId string `json:"shopItemId,omitempty" xml:"shopItemId,omitempty"`
	// 奇门仓储字段,C123,string(50),必填,
	SkuId string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	// 商品编码,S1234,string(50),必填,
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
}

var poolRequestDo = sync.Pool{
	New: func() any {
		return new(RequestDo)
	},
}

// GetRequestDo() 从对象池中获取RequestDo
func GetRequestDo() *RequestDo {
	return poolRequestDo.Get().(*RequestDo)
}

// ReleaseRequestDo 释放RequestDo
func ReleaseRequestDo(v *RequestDo) {
	v.WarehouseCodes = v.WarehouseCodes[:0]
	v.ItemCodes = v.ItemCodes[:0]
	v.ChannelCodes = v.ChannelCodes[:0]
	v.InventoryRules = v.InventoryRules[:0]
	v.OwnerCode = ""
	v.ItemId = ""
	v.ActionType = ""
	v.ShopNick = ""
	v.ItemSource = ""
	v.ShopItemId = ""
	v.SkuId = ""
	v.ItemCode = ""
	poolRequestDo.Put(v)
}
