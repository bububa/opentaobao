package wdk

// AlibabawdkskuchannelskuqueryApiResults 结构体
type AlibabawdkskuchannelskuqueryApiResults struct {
	// 业务数据模型
	Models []AlibabawdkskuchannelskuqueryApiResult `json:"models,omitempty" xml:"models>alibabawdkskuchannelskuquery_api_result,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// skuCode不能为空
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
