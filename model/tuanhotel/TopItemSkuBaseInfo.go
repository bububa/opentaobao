package tuanhotel

import (
	"sync"
)

// TopItemSkuBaseInfo 结构体
type TopItemSkuBaseInfo struct {
	// sku名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolTopItemSkuBaseInfo = sync.Pool{
	New: func() any {
		return new(TopItemSkuBaseInfo)
	},
}

// GetTopItemSkuBaseInfo() 从对象池中获取TopItemSkuBaseInfo
func GetTopItemSkuBaseInfo() *TopItemSkuBaseInfo {
	return poolTopItemSkuBaseInfo.Get().(*TopItemSkuBaseInfo)
}

// ReleaseTopItemSkuBaseInfo 释放TopItemSkuBaseInfo
func ReleaseTopItemSkuBaseInfo(v *TopItemSkuBaseInfo) {
	v.SkuName = ""
	v.OuterId = ""
	v.SkuId = 0
	poolTopItemSkuBaseInfo.Put(v)
}
