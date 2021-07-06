package ieagency

// QueryRefundTicketDetailRs 结构体
type QueryRefundTicketDetailRs struct {
	// apiErrorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// ieRefundTicketVO
	RefundTicket *IeRefundTicketVo `json:"refund_ticket,omitempty" xml:"refund_ticket,omitempty"`
	// apiErrorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
