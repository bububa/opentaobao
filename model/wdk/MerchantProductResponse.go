package wdk

import (
	"sync"
)

// MerchantProductResponse 结构体
type MerchantProductResponse struct {
	// 货品id
	ScIds []int64 `json:"sc_ids,omitempty" xml:"sc_ids>int64,omitempty"`
	// [&#34;123&#34;,&#34;456&#34;]
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolMerchantProductResponse = sync.Pool{
	New: func() any {
		return new(MerchantProductResponse)
	},
}

// GetMerchantProductResponse() 从对象池中获取MerchantProductResponse
func GetMerchantProductResponse() *MerchantProductResponse {
	return poolMerchantProductResponse.Get().(*MerchantProductResponse)
}

// ReleaseMerchantProductResponse 释放MerchantProductResponse
func ReleaseMerchantProductResponse(v *MerchantProductResponse) {
	v.ScIds = v.ScIds[:0]
	v.ItemId = 0
	poolMerchantProductResponse.Put(v)
}
