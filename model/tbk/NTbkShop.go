package tbk

import (
	"sync"
)

// NTbkShop 结构体
type NTbkShop struct {
	// 店铺推广链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 店铺地址
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 店铺类型，B：天猫，C：淘宝
	ShopType string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// 店铺名称
	ShopTitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// 店标图片
	PictUrl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 卖家id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolNTbkShop = sync.Pool{
	New: func() any {
		return new(NTbkShop)
	},
}

// GetNTbkShop() 从对象池中获取NTbkShop
func GetNTbkShop() *NTbkShop {
	return poolNTbkShop.Get().(*NTbkShop)
}

// ReleaseNTbkShop 释放NTbkShop
func ReleaseNTbkShop(v *NTbkShop) {
	v.ClickUrl = ""
	v.ShopUrl = ""
	v.SellerNick = ""
	v.ShopType = ""
	v.ShopTitle = ""
	v.PictUrl = ""
	v.UserId = 0
	poolNTbkShop.Put(v)
}
