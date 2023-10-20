package wdk

import (
	"sync"
)

// PreCheckResponse 结构体
type PreCheckResponse struct {
	// 商品是否可作业信息
	ProductList []PreCheckProductResponse `json:"product_list,omitempty" xml:"product_list>pre_check_product_response,omitempty"`
	// 履约前置校验扩展数据
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
}

var poolPreCheckResponse = sync.Pool{
	New: func() any {
		return new(PreCheckResponse)
	},
}

// GetPreCheckResponse() 从对象池中获取PreCheckResponse
func GetPreCheckResponse() *PreCheckResponse {
	return poolPreCheckResponse.Get().(*PreCheckResponse)
}

// ReleasePreCheckResponse 释放PreCheckResponse
func ReleasePreCheckResponse(v *PreCheckResponse) {
	v.ProductList = v.ProductList[:0]
	v.Ext = ""
	poolPreCheckResponse.Put(v)
}
