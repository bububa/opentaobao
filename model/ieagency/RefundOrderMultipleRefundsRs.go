package ieagency

// RefundOrderMultipleRefundsRs 
type RefundOrderMultipleRefundsRs struct {
    // 错误码
    ApiErrorCode   int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
    // 错误文案
    ApiErrorMsg   string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
    // 错误详情
    ErrTrace   string `json:"err_trace,omitempty" xml:"err_trace,omitempty"`
    // 失败
    Failure   bool `json:"failure,omitempty" xml:"failure,omitempty"`
    // hostName
    HostName   string `json:"host_name,omitempty" xml:"host_name,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // refundTicketId
    RefundTicketId   int64 `json:"refund_ticket_id,omitempty" xml:"refund_ticket_id,omitempty"`
}
