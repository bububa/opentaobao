package wdk

import (
	"sync"
)

// PreCheckProductResponse 结构体
type PreCheckProductResponse struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 是否可作业
	CanFulfill bool `json:"can_fulfill,omitempty" xml:"can_fulfill,omitempty"`
}

var poolPreCheckProductResponse = sync.Pool{
	New: func() any {
		return new(PreCheckProductResponse)
	},
}

// GetPreCheckProductResponse() 从对象池中获取PreCheckProductResponse
func GetPreCheckProductResponse() *PreCheckProductResponse {
	return poolPreCheckProductResponse.Get().(*PreCheckProductResponse)
}

// ReleasePreCheckProductResponse 释放PreCheckProductResponse
func ReleasePreCheckProductResponse(v *PreCheckProductResponse) {
	v.SkuCode = ""
	v.CanFulfill = false
	poolPreCheckProductResponse.Put(v)
}
