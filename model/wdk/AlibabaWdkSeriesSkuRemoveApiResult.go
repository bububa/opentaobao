package wdk

// AlibabaWdkSeriesSkuRemoveApiResult 结构体
type AlibabaWdkSeriesSkuRemoveApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误详情
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 系列品移除商品成功结果
	Model *SkuSeriesEditResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
