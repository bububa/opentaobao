package trade

// TaobaoRdcAligeniusOrdermsgUpdateResult 结构体
type TaobaoRdcAligeniusOrdermsgUpdateResult struct {
	// resultData
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// errorInfo
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
