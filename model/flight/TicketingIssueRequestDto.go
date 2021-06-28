package flight

// TicketingIssueRequestDto 
/* model for simplify = false
type TicketingIssueRequestDto struct {

    // 国内国际标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`
    

    // 飞猪订单号
    
    OrderId   string `json:"order_id,omitempty"`
    

    // 出票信息
    
    IssueList  struct {
        TicketingPsgItemDto  []TicketingPsgItemDto `json:"ticketing_psg_item_dto,omitempty"`
    } `json:"issue_list,omitempty"`
    

}
*/

// TicketingIssueRequestDto 
type TicketingIssueRequestDto struct {

    // 国内国际标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`

    // 飞猪订单号
    OrderId   string `json:"order_id,omitempty"`

    // 出票信息
    IssueList   []TicketingPsgItemDto `json:"issue_list,omitempty"`

}
