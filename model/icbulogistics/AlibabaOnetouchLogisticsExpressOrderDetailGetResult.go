package icbulogistics

// AlibabaonetouchlogisticsexpressorderdetailgetResult 结构体
type AlibabaonetouchlogisticsexpressorderdetailgetResult struct {
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果数据
	Data *OrderDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
