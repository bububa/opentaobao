package ascp

import (
	"sync"
)

// OrderProcessReportRequest 结构体
type OrderProcessReportRequest struct {
	// 单据行
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 货主ID
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 单据信息
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
	// 单据作业信息
	Process *Process `json:"process,omitempty" xml:"process,omitempty"`
	// 业务请求时间（时间戳）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolOrderProcessReportRequest = sync.Pool{
	New: func() any {
		return new(OrderProcessReportRequest)
	},
}

// GetOrderProcessReportRequest() 从对象池中获取OrderProcessReportRequest
func GetOrderProcessReportRequest() *OrderProcessReportRequest {
	return poolOrderProcessReportRequest.Get().(*OrderProcessReportRequest)
}

// ReleaseOrderProcessReportRequest 释放OrderProcessReportRequest
func ReleaseOrderProcessReportRequest(v *OrderProcessReportRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.RequestId = ""
	v.ExtendProps = ""
	v.OwnerCode = ""
	v.Order = nil
	v.Process = nil
	v.RequestTime = 0
	poolOrderProcessReportRequest.Put(v)
}
