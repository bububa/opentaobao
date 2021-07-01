package icbulogistics

// AlibabaOnetouchLogisticsExpressOrderDetailGetResult 结构体
type AlibabaOnetouchLogisticsExpressOrderDetailGetResult struct {
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 结果数据
	Data *OrderDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}
