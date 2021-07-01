package wdk

// MaochaoOrderInsuranceCallbackResult 结构体
type MaochaoOrderInsuranceCallbackResult struct {
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
