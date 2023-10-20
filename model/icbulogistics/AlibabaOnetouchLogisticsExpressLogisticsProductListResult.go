package icbulogistics

// AlibabaonetouchlogisticsexpresslogisticsproductlistResult 结构体
type AlibabaonetouchlogisticsexpresslogisticsproductlistResult struct {
	// 列表对象
	Values []LogisticsProductDto `json:"values,omitempty" xml:"values>logistics_product_dto,omitempty"`
	// 错误信息提示
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
