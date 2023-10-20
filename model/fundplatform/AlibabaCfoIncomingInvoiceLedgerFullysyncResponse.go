package fundplatform

// AlibabaCfoIncomingInvoiceLedgerFullysyncResponse 结构体
type AlibabaCfoIncomingInvoiceLedgerFullysyncResponse struct {
	// 响应码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 响应消息
	ResponseMsg string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
	// 是否成功
	Succeeded string `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
