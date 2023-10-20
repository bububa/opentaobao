package iotticket

// CainiaoiotticketdetailqueryResult 结构体
type CainiaoiotticketdetailqueryResult struct {
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果数据
	Data *CainiaoiotticketdetailqueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
