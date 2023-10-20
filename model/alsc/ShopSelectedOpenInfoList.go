package alsc

import (
	"sync"
)

// ShopSelectedOpenInfoList 结构体
type ShopSelectedOpenInfoList struct {
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
}

var poolShopSelectedOpenInfoList = sync.Pool{
	New: func() any {
		return new(ShopSelectedOpenInfoList)
	},
}

// GetShopSelectedOpenInfoList() 从对象池中获取ShopSelectedOpenInfoList
func GetShopSelectedOpenInfoList() *ShopSelectedOpenInfoList {
	return poolShopSelectedOpenInfoList.Get().(*ShopSelectedOpenInfoList)
}

// ReleaseShopSelectedOpenInfoList 释放ShopSelectedOpenInfoList
func ReleaseShopSelectedOpenInfoList(v *ShopSelectedOpenInfoList) {
	v.ShopId = ""
	v.OutShopId = ""
	poolShopSelectedOpenInfoList.Put(v)
}
