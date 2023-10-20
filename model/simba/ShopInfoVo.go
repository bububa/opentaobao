package simba

import (
	"sync"
)

// ShopInfoVo 结构体
type ShopInfoVo struct {
	// 店铺名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
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
	v.Title = ""
	v.ShopId = 0
	poolShopInfoVo.Put(v)
}
