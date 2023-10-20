package tmallsc

import (
	"sync"
)

// ServicePriceQueryResponse 结构体
type ServicePriceQueryResponse struct {
	// 服务价格记录行
	ServicePriceLines []ServicePriceLine `json:"service_price_lines,omitempty" xml:"service_price_lines>service_price_line,omitempty"`
}

var poolServicePriceQueryResponse = sync.Pool{
	New: func() any {
		return new(ServicePriceQueryResponse)
	},
}

// GetServicePriceQueryResponse() 从对象池中获取ServicePriceQueryResponse
func GetServicePriceQueryResponse() *ServicePriceQueryResponse {
	return poolServicePriceQueryResponse.Get().(*ServicePriceQueryResponse)
}

// ReleaseServicePriceQueryResponse 释放ServicePriceQueryResponse
func ReleaseServicePriceQueryResponse(v *ServicePriceQueryResponse) {
	v.ServicePriceLines = v.ServicePriceLines[:0]
	poolServicePriceQueryResponse.Put(v)
}
