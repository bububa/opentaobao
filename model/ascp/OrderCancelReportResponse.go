package ascp

import (
	"sync"
)

// OrderCancelReportResponse 结构体
type OrderCancelReportResponse struct {
	// traceId，类似于requestId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderCancelReportResponse = sync.Pool{
	New: func() any {
		return new(OrderCancelReportResponse)
	},
}

// GetOrderCancelReportResponse() 从对象池中获取OrderCancelReportResponse
func GetOrderCancelReportResponse() *OrderCancelReportResponse {
	return poolOrderCancelReportResponse.Get().(*OrderCancelReportResponse)
}

// ReleaseOrderCancelReportResponse 释放OrderCancelReportResponse
func ReleaseOrderCancelReportResponse(v *OrderCancelReportResponse) {
	v.TraceId = ""
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolOrderCancelReportResponse.Put(v)
}
