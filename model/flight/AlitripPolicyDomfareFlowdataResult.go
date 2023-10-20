package flight

// AlitripPolicyDomfareFlowdataResult 结构体
type AlitripPolicyDomfareFlowdataResult struct {
	// 返回错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回的错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回的信息
	Data *CompareFlowDataReponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
