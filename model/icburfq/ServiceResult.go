package icburfq

// ServiceResult 结构体
type ServiceResult struct {
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 我的权益信息
	Value *EquityPackageDto `json:"value,omitempty" xml:"value,omitempty"`
	// 操作结果对象
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
