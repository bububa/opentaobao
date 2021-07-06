package ieagency

// RefundOrderMultipleRefundsRs 结构体
type RefundOrderMultipleRefundsRs struct {
	// 错误文案
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 错误详情
	ErrTrace string `json:"err_trace,omitempty" xml:"err_trace,omitempty"`
	// hostName
	HostName string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// 错误码
	ApiErrorCode int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// refundTicketId
	RefundTicketId int64 `json:"refund_ticket_id,omitempty" xml:"refund_ticket_id,omitempty"`
	// 失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
