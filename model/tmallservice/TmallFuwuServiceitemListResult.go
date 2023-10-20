package tmallservice

// TmallfuwuserviceitemlistResult 结构体
type TmallfuwuserviceitemlistResult struct {
	// 服务商品信息列表的json对象
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorType
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
