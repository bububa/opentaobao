package logistic

// TaobaonextonelogisticssignupdateResult 结构体
type TaobaonextonelogisticssignupdateResult struct {
	// 返回数据
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 错误信息
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
