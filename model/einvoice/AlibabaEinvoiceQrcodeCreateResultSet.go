package einvoice

// AlibabaEinvoiceQrcodeCreateResultSet 结构体
type AlibabaEinvoiceQrcodeCreateResultSet struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
