package alsc

import (
	"sync"
)

// ShopOpenInfo 结构体
type ShopOpenInfo struct {
	// 店铺扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 店铺Pid
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 店铺ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 店铺logo图片
	ShopLogoUrl string `json:"shop_logo_url,omitempty" xml:"shop_logo_url,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
}

var poolShopOpenInfo = sync.Pool{
	New: func() any {
		return new(ShopOpenInfo)
	},
}

// GetShopOpenInfo() 从对象池中获取ShopOpenInfo
func GetShopOpenInfo() *ShopOpenInfo {
	return poolShopOpenInfo.Get().(*ShopOpenInfo)
}

// ReleaseShopOpenInfo 释放ShopOpenInfo
func ReleaseShopOpenInfo(v *ShopOpenInfo) {
	v.ExtInfo = ""
	v.Pid = ""
	v.ShopId = ""
	v.ShopLogoUrl = ""
	v.ShopName = ""
	poolShopOpenInfo.Put(v)
}
