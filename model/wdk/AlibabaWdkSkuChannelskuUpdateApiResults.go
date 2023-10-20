package wdk

// AlibabawdkskuchannelskuupdateApiResults 结构体
type AlibabawdkskuchannelskuupdateApiResults struct {
	// 单个商品返回结果集合
	Models []AlibabawdkskuchannelskuupdateApiResult `json:"models,omitempty" xml:"models>alibabawdkskuchannelskuupdate_api_result,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
