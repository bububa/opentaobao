package ascp

import (
	"sync"
)

// OrderProcessReportReponse 结构体
type OrderProcessReportReponse struct {
	// traceId，类似于requestId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderProcessReportReponse = sync.Pool{
	New: func() any {
		return new(OrderProcessReportReponse)
	},
}

// GetOrderProcessReportReponse() 从对象池中获取OrderProcessReportReponse
func GetOrderProcessReportReponse() *OrderProcessReportReponse {
	return poolOrderProcessReportReponse.Get().(*OrderProcessReportReponse)
}

// ReleaseOrderProcessReportReponse 释放OrderProcessReportReponse
func ReleaseOrderProcessReportReponse(v *OrderProcessReportReponse) {
	v.TraceId = ""
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolOrderProcessReportReponse.Put(v)
}
