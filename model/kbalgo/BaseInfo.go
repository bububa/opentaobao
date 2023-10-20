package kbalgo

import (
	"sync"
)

// BaseInfo 结构体
type BaseInfo struct {
	// 是否有优惠
	ShopPromotion string `json:"shop_promotion,omitempty" xml:"shop_promotion,omitempty"`
}

var poolBaseInfo = sync.Pool{
	New: func() any {
		return new(BaseInfo)
	},
}

// GetBaseInfo() 从对象池中获取BaseInfo
func GetBaseInfo() *BaseInfo {
	return poolBaseInfo.Get().(*BaseInfo)
}

// ReleaseBaseInfo 释放BaseInfo
func ReleaseBaseInfo(v *BaseInfo) {
	v.ShopPromotion = ""
	poolBaseInfo.Put(v)
}
