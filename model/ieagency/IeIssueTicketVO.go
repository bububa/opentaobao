package ieagency

// IeIssueTicketVO 
type IeIssueTicketVO struct {

    // 预定订单pnr信息
    
    UpdatePnrVos   []IeBookPnrVo `json:"update_pnr_vos,omitempty" xml:"update_pnr_vos,omitempty"`
    

    // 乘机人票信息
    
    PassengerTicketVos   []IePassengerTicketVO `json:"passenger_ticket_vos,omitempty" xml:"passenger_ticket_vos,omitempty"`
    

    // 预定订单id
    
    BookOrderId   int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
    

    // 订单备注
    
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    

}
