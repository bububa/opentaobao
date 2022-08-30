package ascp

// BaseResult 结构体
type BaseResult struct {
	// 通常用于success为false时的页面错误类型判定
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 通常用于success为false时的页面错误信息提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 通常用于success为true时的页面提示
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 数据
	Result *HiErpOrderResp `json:"result,omitempty" xml:"result,omitempty"`
	// false: 失败   true: 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
