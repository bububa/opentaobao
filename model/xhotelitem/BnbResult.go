package xhotelitem

// BnbResult 结构体
type BnbResult struct {
	// 响应码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 响应信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 状态，成功true，失败false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
