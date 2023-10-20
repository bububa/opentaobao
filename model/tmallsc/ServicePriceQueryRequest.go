package tmallsc

import (
	"sync"
)

// ServicePriceQueryRequest 结构体
type ServicePriceQueryRequest struct {
	// 服务外部编码
	ServiceOuterIds []string `json:"service_outer_ids,omitempty" xml:"service_outer_ids>string,omitempty"`
	// 卖家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolServicePriceQueryRequest = sync.Pool{
	New: func() any {
		return new(ServicePriceQueryRequest)
	},
}

// GetServicePriceQueryRequest() 从对象池中获取ServicePriceQueryRequest
func GetServicePriceQueryRequest() *ServicePriceQueryRequest {
	return poolServicePriceQueryRequest.Get().(*ServicePriceQueryRequest)
}

// ReleaseServicePriceQueryRequest 释放ServicePriceQueryRequest
func ReleaseServicePriceQueryRequest(v *ServicePriceQueryRequest) {
	v.ServiceOuterIds = v.ServiceOuterIds[:0]
	v.SellerId = ""
	poolServicePriceQueryRequest.Put(v)
}
