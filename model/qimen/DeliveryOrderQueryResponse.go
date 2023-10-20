package qimen

import (
	"sync"
)

// DeliveryOrderQueryResponse 结构体
type DeliveryOrderQueryResponse struct {
	// 包裹信息
	Packages []Package `json:"packages,omitempty" xml:"packages>package,omitempty"`
	// 单据列表
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// orderLines总条数
	TotalLines int64 `json:"totalLines,omitempty" xml:"totalLines,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
}

var poolDeliveryOrderQueryResponse = sync.Pool{
	New: func() any {
		return new(DeliveryOrderQueryResponse)
	},
}

// GetDeliveryOrderQueryResponse() 从对象池中获取DeliveryOrderQueryResponse
func GetDeliveryOrderQueryResponse() *DeliveryOrderQueryResponse {
	return poolDeliveryOrderQueryResponse.Get().(*DeliveryOrderQueryResponse)
}

// ReleaseDeliveryOrderQueryResponse 释放DeliveryOrderQueryResponse
func ReleaseDeliveryOrderQueryResponse(v *DeliveryOrderQueryResponse) {
	v.Packages = v.Packages[:0]
	v.OrderLines = v.OrderLines[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.TotalLines = 0
	v.DeliveryOrder = nil
	poolDeliveryOrderQueryResponse.Put(v)
}
