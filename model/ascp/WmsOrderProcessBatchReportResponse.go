package ascp

import (
	"sync"
)

// WmsOrderProcessBatchReportResponse 结构体
type WmsOrderProcessBatchReportResponse struct {
	// 详细信息
	Data []DataDetails `json:"data,omitempty" xml:"data>data_details,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// traceId，用于排查问题
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWmsOrderProcessBatchReportResponse = sync.Pool{
	New: func() any {
		return new(WmsOrderProcessBatchReportResponse)
	},
}

// GetWmsOrderProcessBatchReportResponse() 从对象池中获取WmsOrderProcessBatchReportResponse
func GetWmsOrderProcessBatchReportResponse() *WmsOrderProcessBatchReportResponse {
	return poolWmsOrderProcessBatchReportResponse.Get().(*WmsOrderProcessBatchReportResponse)
}

// ReleaseWmsOrderProcessBatchReportResponse 释放WmsOrderProcessBatchReportResponse
func ReleaseWmsOrderProcessBatchReportResponse(v *WmsOrderProcessBatchReportResponse) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.TraceId = ""
	v.Success = false
	poolWmsOrderProcessBatchReportResponse.Put(v)
}
