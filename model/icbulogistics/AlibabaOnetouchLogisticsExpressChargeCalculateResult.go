package icbulogistics

// AlibabaOnetouchLogisticsExpressChargeCalculateResult 结构体
type AlibabaOnetouchLogisticsExpressChargeCalculateResult struct {
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 结果对象
	Values *ExpressFreightDto `json:"values,omitempty" xml:"values,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
