package wdk

// AlibabatclsaxintegrationaccountimportApiResult 结构体
type AlibabatclsaxintegrationaccountimportApiResult struct {
	// 扩展信息
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 导入结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
