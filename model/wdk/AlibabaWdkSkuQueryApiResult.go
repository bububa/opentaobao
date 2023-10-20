package wdk

// AlibabaWdkSkuQueryApiResult 结构体
type AlibabaWdkSkuQueryApiResult struct {
	// 请求参数不能为空
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 单条错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 单条查询结果
	Model *SkuDo `json:"model,omitempty" xml:"model,omitempty"`
	// 单条是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
