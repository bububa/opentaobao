package xhotelonlineorder

// HbsResult 结构体
type HbsResult struct {
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 给开发用的错误信息
	ResultMsg4Dev string `json:"result_msg4_dev,omitempty" xml:"result_msg4_dev,omitempty"`
	// 响应扩展信息
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 泛型结果
	Module *OutSourceOrderCreateRes `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
