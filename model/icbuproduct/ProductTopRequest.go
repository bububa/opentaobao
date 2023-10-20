package icbuproduct

import (
	"sync"
)

// ProductTopRequest 结构体
type ProductTopRequest struct {
	// 返回结果语种
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}

var poolProductTopRequest = sync.Pool{
	New: func() any {
		return new(ProductTopRequest)
	},
}

// GetProductTopRequest() 从对象池中获取ProductTopRequest
func GetProductTopRequest() *ProductTopRequest {
	return poolProductTopRequest.Get().(*ProductTopRequest)
}

// ReleaseProductTopRequest 释放ProductTopRequest
func ReleaseProductTopRequest(v *ProductTopRequest) {
	v.Language = ""
	v.CatId = 0
	poolProductTopRequest.Put(v)
}
