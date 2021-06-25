package bus

// B2BFetchTicket 
type B2BFetchTicket struct {

    // 取票短信
    FetchSms   string `json:"fetch_sms,omitempty"`

    // 取票地址
    FetchTicketAddress   string `json:"fetch_ticket_address,omitempty"`

    // 取票号
    FetchTicketNumber   string `json:"fetch_ticket_number,omitempty"`

    // 取票密码
    FetchTicketPwd   string `json:"fetch_ticket_pwd,omitempty"`

}
