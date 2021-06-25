package bus

// B2BQueryRefundFeeRp 
type B2BQueryRefundFeeRp struct {

    // canReturnSingleTicket
    CanReturnSingleTicket   bool `json:"can_return_single_ticket,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // refundFees
    RefundFees   []Json `json:"refund_fees,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
