package promotion

import (
	"sync"
)

// ShopInfoVo 结构体
type ShopInfoVo struct {
	// 店铺链接
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// 店铺icon链接
	ShopIconUrl string `json:"shop_icon_url,omitempty" xml:"shop_icon_url,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 店铺id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolShopInfoVo = sync.Pool{
	New: func() any {
		return new(ShopInfoVo)
	},
}

// GetShopInfoVo() 从对象池中获取ShopInfoVo
func GetShopInfoVo() *ShopInfoVo {
	return poolShopInfoVo.Get().(*ShopInfoVo)
}

// ReleaseShopInfoVo 释放ShopInfoVo
func ReleaseShopInfoVo(v *ShopInfoVo) {
	v.ShopUrl = ""
	v.ShopIconUrl = ""
	v.ShopName = ""
	v.SellerId = 0
	v.ShopId = 0
	poolShopInfoVo.Put(v)
}
