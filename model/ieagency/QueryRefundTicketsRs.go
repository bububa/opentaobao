package ieagency

// QueryRefundTicketsRs 
type QueryRefundTicketsRs struct {

    // refundTicketList
    
    RefundTicketList   []IeRefundTicketVo `json:"refund_ticket_list,omitempty" xml:"refund_ticket_list,omitempty"`
    

    // apiErrorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // apiErrorCode
    
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // pageCount
    
    PageCount   int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
    

}
