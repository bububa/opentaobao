package ascp

import (
	"sync"
)

// OrderProcessReportResponse 结构体
type OrderProcessReportResponse struct {
	// traceId，类似于requestId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderProcessReportResponse = sync.Pool{
	New: func() any {
		return new(OrderProcessReportResponse)
	},
}

// GetOrderProcessReportResponse() 从对象池中获取OrderProcessReportResponse
func GetOrderProcessReportResponse() *OrderProcessReportResponse {
	return poolOrderProcessReportResponse.Get().(*OrderProcessReportResponse)
}

// ReleaseOrderProcessReportResponse 释放OrderProcessReportResponse
func ReleaseOrderProcessReportResponse(v *OrderProcessReportResponse) {
	v.TraceId = ""
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolOrderProcessReportResponse.Put(v)
}
