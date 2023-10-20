package examination

// AlibabaalihealthmedicalorderrefundResult 结构体
type AlibabaalihealthmedicalorderrefundResult struct {
	// SUCCESS:成功; FAIL:失败; UNKNOWN:未知;
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回数据
	Data *OrderRefundVo `json:"data,omitempty" xml:"data,omitempty"`
}
