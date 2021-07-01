package dt

// NrsResult 结构体
type NrsResult struct {
	// 返回数据
	Data *RecongnizeItemInfo `json:"data,omitempty" xml:"data,omitempty"`
	// 接口调用标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
