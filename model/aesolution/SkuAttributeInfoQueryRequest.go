package aesolution

import (
	"sync"
)

// SkuAttributeInfoQueryRequest 结构体
type SkuAttributeInfoQueryRequest struct {
	// merchant&#39;s category ID
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// aliexpress category ID. aliexpress_category_id and category_id could not be both empty.
	AliexpressCategoryId int64 `json:"aliexpress_category_id,omitempty" xml:"aliexpress_category_id,omitempty"`
}

var poolSkuAttributeInfoQueryRequest = sync.Pool{
	New: func() any {
		return new(SkuAttributeInfoQueryRequest)
	},
}

// GetSkuAttributeInfoQueryRequest() 从对象池中获取SkuAttributeInfoQueryRequest
func GetSkuAttributeInfoQueryRequest() *SkuAttributeInfoQueryRequest {
	return poolSkuAttributeInfoQueryRequest.Get().(*SkuAttributeInfoQueryRequest)
}

// ReleaseSkuAttributeInfoQueryRequest 释放SkuAttributeInfoQueryRequest
func ReleaseSkuAttributeInfoQueryRequest(v *SkuAttributeInfoQueryRequest) {
	v.CategoryId = ""
	v.AliexpressCategoryId = 0
	poolSkuAttributeInfoQueryRequest.Put(v)
}
