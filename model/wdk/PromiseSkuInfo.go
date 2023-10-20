package wdk

import (
	"sync"
)

// PromiseSkuInfo 结构体
type PromiseSkuInfo struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品履约线路
	LineInstances string `json:"line_instances,omitempty" xml:"line_instances,omitempty"`
	// 加购数量
	Quantity float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolPromiseSkuInfo = sync.Pool{
	New: func() any {
		return new(PromiseSkuInfo)
	},
}

// GetPromiseSkuInfo() 从对象池中获取PromiseSkuInfo
func GetPromiseSkuInfo() *PromiseSkuInfo {
	return poolPromiseSkuInfo.Get().(*PromiseSkuInfo)
}

// ReleasePromiseSkuInfo 释放PromiseSkuInfo
func ReleasePromiseSkuInfo(v *PromiseSkuInfo) {
	v.SkuCode = ""
	v.LineInstances = ""
	v.Quantity = 0
	poolPromiseSkuInfo.Put(v)
}
