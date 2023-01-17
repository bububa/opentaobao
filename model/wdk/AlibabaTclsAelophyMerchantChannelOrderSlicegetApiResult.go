package wdk

// AlibabaTclsAelophyMerchantChannelOrderSlicegetApiResult 结构体
type AlibabaTclsAelophyMerchantChannelOrderSlicegetApiResult struct {
	// 返回码说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回的时间片数据
	Model *TimeSliceGetResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
