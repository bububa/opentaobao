package ieagency

// ReceiveRefundTicketRs 
type ReceiveRefundTicketRs struct {
    // apiErrorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // apiErrorCode
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}