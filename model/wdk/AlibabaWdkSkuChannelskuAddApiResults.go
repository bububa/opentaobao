package wdk

// AlibabawdkskuchannelskuaddApiResults 结构体
type AlibabawdkskuchannelskuaddApiResults struct {
	// 返会结果集合
	Models []AlibabawdkskuchannelskuaddApiResult `json:"models,omitempty" xml:"models>alibabawdkskuchannelskuadd_api_result,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
