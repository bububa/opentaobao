package wdk

// AlibabawdkskuupdateApiResults 结构体
type AlibabawdkskuupdateApiResults struct {
	// 各条记录结果
	Models []AlibabawdkskuupdateApiResult `json:"models,omitempty" xml:"models>alibabawdkskuupdate_api_result,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口调用成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
