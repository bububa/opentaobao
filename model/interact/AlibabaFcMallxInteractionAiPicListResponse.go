package interact

// AlibabaFcMallxInteractionAiPicListResponse 结构体
type AlibabaFcMallxInteractionAiPicListResponse struct {
	// 结果code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果信息
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 结果数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
