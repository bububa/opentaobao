package alsc

import (
	"sync"
)

// ShopSelectedOpenInfo 结构体
type ShopSelectedOpenInfo struct {
	// 门店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部门店ID
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
}

var poolShopSelectedOpenInfo = sync.Pool{
	New: func() any {
		return new(ShopSelectedOpenInfo)
	},
}

// GetShopSelectedOpenInfo() 从对象池中获取ShopSelectedOpenInfo
func GetShopSelectedOpenInfo() *ShopSelectedOpenInfo {
	return poolShopSelectedOpenInfo.Get().(*ShopSelectedOpenInfo)
}

// ReleaseShopSelectedOpenInfo 释放ShopSelectedOpenInfo
func ReleaseShopSelectedOpenInfo(v *ShopSelectedOpenInfo) {
	v.ShopId = ""
	v.OutShopId = ""
	poolShopSelectedOpenInfo.Put(v)
}
