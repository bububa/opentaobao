package bus

// B2BQueryRefundFeeRp 结构体
type B2BQueryRefundFeeRp struct {
	// refundFees
	RefundFees []string `json:"refund_fees,omitempty" xml:"refund_fees>string,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// canReturnSingleTicket
	CanReturnSingleTicket bool `json:"can_return_single_ticket,omitempty" xml:"can_return_single_ticket,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
