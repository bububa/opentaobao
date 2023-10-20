package ascp

import (
	"sync"
)

// DeliveryOrderConfirmReportResponse 结构体
type DeliveryOrderConfirmReportResponse struct {
	// traceId，类似于requestId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeliveryOrderConfirmReportResponse = sync.Pool{
	New: func() any {
		return new(DeliveryOrderConfirmReportResponse)
	},
}

// GetDeliveryOrderConfirmReportResponse() 从对象池中获取DeliveryOrderConfirmReportResponse
func GetDeliveryOrderConfirmReportResponse() *DeliveryOrderConfirmReportResponse {
	return poolDeliveryOrderConfirmReportResponse.Get().(*DeliveryOrderConfirmReportResponse)
}

// ReleaseDeliveryOrderConfirmReportResponse 释放DeliveryOrderConfirmReportResponse
func ReleaseDeliveryOrderConfirmReportResponse(v *DeliveryOrderConfirmReportResponse) {
	v.TraceId = ""
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolDeliveryOrderConfirmReportResponse.Put(v)
}
