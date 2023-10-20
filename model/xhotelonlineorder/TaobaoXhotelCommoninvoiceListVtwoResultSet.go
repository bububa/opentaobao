package xhotelonlineorder

// TaobaoxhotelcommoninvoicelistvtwoResultSet 结构体
type TaobaoxhotelcommoninvoicelistvtwoResultSet struct {
	// 常用发票信息
	Results []CommonInvoiceInfo `json:"results,omitempty" xml:"results>common_invoice_info,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
