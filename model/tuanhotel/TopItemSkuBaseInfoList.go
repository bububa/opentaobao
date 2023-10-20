package tuanhotel

import (
	"sync"
)

// TopItemSkuBaseInfoList 结构体
type TopItemSkuBaseInfoList struct {
	// sku名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolTopItemSkuBaseInfoList = sync.Pool{
	New: func() any {
		return new(TopItemSkuBaseInfoList)
	},
}

// GetTopItemSkuBaseInfoList() 从对象池中获取TopItemSkuBaseInfoList
func GetTopItemSkuBaseInfoList() *TopItemSkuBaseInfoList {
	return poolTopItemSkuBaseInfoList.Get().(*TopItemSkuBaseInfoList)
}

// ReleaseTopItemSkuBaseInfoList 释放TopItemSkuBaseInfoList
func ReleaseTopItemSkuBaseInfoList(v *TopItemSkuBaseInfoList) {
	v.SkuName = ""
	v.OuterId = ""
	v.SkuId = 0
	poolTopItemSkuBaseInfoList.Put(v)
}
