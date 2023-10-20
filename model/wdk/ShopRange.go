package wdk

import (
	"sync"
)

// ShopRange 结构体
type ShopRange struct {
	// 经纬度点
	Points []Point `json:"points,omitempty" xml:"points>point,omitempty"`
}

var poolShopRange = sync.Pool{
	New: func() any {
		return new(ShopRange)
	},
}

// GetShopRange() 从对象池中获取ShopRange
func GetShopRange() *ShopRange {
	return poolShopRange.Get().(*ShopRange)
}

// ReleaseShopRange 释放ShopRange
func ReleaseShopRange(v *ShopRange) {
	v.Points = v.Points[:0]
	poolShopRange.Put(v)
}
