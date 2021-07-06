package yunos

// RdamResponse 结构体
type RdamResponse struct {
	// dataList
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// more
	More string `json:"more,omitempty" xml:"more,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
