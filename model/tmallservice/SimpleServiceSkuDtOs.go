package tmallservice

import (
	"sync"
)

// SimpleServiceSkuDtOs 结构体
type SimpleServiceSkuDtOs struct {
	// 服务skuCode
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 服务sku名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
}

var poolSimpleServiceSkuDtOs = sync.Pool{
	New: func() any {
		return new(SimpleServiceSkuDtOs)
	},
}

// GetSimpleServiceSkuDtOs() 从对象池中获取SimpleServiceSkuDtOs
func GetSimpleServiceSkuDtOs() *SimpleServiceSkuDtOs {
	return poolSimpleServiceSkuDtOs.Get().(*SimpleServiceSkuDtOs)
}

// ReleaseSimpleServiceSkuDtOs 释放SimpleServiceSkuDtOs
func ReleaseSimpleServiceSkuDtOs(v *SimpleServiceSkuDtOs) {
	v.Code = ""
	v.DisplayName = ""
	poolSimpleServiceSkuDtOs.Put(v)
}
