package ascp

// WmsOrderProcessReportResponse 结构体
type WmsOrderProcessReportResponse struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// traceId，类似于requestId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
