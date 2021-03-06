package retail

// Detailresults 结构体
type Detailresults struct {
	// outOrderId
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
