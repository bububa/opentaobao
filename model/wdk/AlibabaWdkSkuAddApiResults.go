package wdk

// AlibabaWdkSkuAddApiResults 结构体
type AlibabaWdkSkuAddApiResults struct {
	// models
	Models []AlibabaWdkSkuAddApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_add_api_result,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口返回成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
