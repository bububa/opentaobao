package icbudropshipping

import (
	"sync"
)

// DistributionSaleProductRequest 结构体
type DistributionSaleProductRequest struct {
	// productId List，max size is 10
	ProductIds []int64 `json:"product_ids,omitempty" xml:"product_ids>int64,omitempty"`
}

var poolDistributionSaleProductRequest = sync.Pool{
	New: func() any {
		return new(DistributionSaleProductRequest)
	},
}

// GetDistributionSaleProductRequest() 从对象池中获取DistributionSaleProductRequest
func GetDistributionSaleProductRequest() *DistributionSaleProductRequest {
	return poolDistributionSaleProductRequest.Get().(*DistributionSaleProductRequest)
}

// ReleaseDistributionSaleProductRequest 释放DistributionSaleProductRequest
func ReleaseDistributionSaleProductRequest(v *DistributionSaleProductRequest) {
	v.ProductIds = v.ProductIds[:0]
	poolDistributionSaleProductRequest.Put(v)
}
