package ascp

import (
	"sync"
)

// DeliveryOrderReportResponse 结构体
type DeliveryOrderReportResponse struct {
	// traceId，类似于requestId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeliveryOrderReportResponse = sync.Pool{
	New: func() any {
		return new(DeliveryOrderReportResponse)
	},
}

// GetDeliveryOrderReportResponse() 从对象池中获取DeliveryOrderReportResponse
func GetDeliveryOrderReportResponse() *DeliveryOrderReportResponse {
	return poolDeliveryOrderReportResponse.Get().(*DeliveryOrderReportResponse)
}

// ReleaseDeliveryOrderReportResponse 释放DeliveryOrderReportResponse
func ReleaseDeliveryOrderReportResponse(v *DeliveryOrderReportResponse) {
	v.TraceId = ""
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolDeliveryOrderReportResponse.Put(v)
}
