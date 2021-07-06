package user

// TaobaoRdcAligeniusAccountValidateResult 结构体
type TaobaoRdcAligeniusAccountValidateResult struct {
	// errorInfo
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// resultData
	ResultData *Resultdata `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 为true时才有resultData
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
