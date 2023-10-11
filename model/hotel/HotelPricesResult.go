package hotel

// HotelPricesResult 结构体
type HotelPricesResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码，不为0表示有异常
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 酒店报价信息
	Module *HotelPricesResponse `json:"module,omitempty" xml:"module,omitempty"`
}
