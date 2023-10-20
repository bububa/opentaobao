package ascp

import (
	"sync"
)

// DeliveryOrderReportRequest 结构体
type DeliveryOrderReportRequest struct {
	// 订单列表
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 货主ID
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 单据信息
	DeliveryOrder *DeliveryOrder `json:"delivery_order,omitempty" xml:"delivery_order,omitempty"`
	// 业务请求时间（时间戳）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolDeliveryOrderReportRequest = sync.Pool{
	New: func() any {
		return new(DeliveryOrderReportRequest)
	},
}

// GetDeliveryOrderReportRequest() 从对象池中获取DeliveryOrderReportRequest
func GetDeliveryOrderReportRequest() *DeliveryOrderReportRequest {
	return poolDeliveryOrderReportRequest.Get().(*DeliveryOrderReportRequest)
}

// ReleaseDeliveryOrderReportRequest 释放DeliveryOrderReportRequest
func ReleaseDeliveryOrderReportRequest(v *DeliveryOrderReportRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.ExtendProps = ""
	v.RequestId = ""
	v.OwnerCode = ""
	v.DeliveryOrder = nil
	v.RequestTime = 0
	poolDeliveryOrderReportRequest.Put(v)
}
