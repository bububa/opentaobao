package bus

// B2BFetchHolderInfo 
/* model for simplify = false
type B2BFetchHolderInfo struct {

    // 取票人证件号码
    
    FetchCertNumber   string `json:"fetch_cert_number,omitempty"`
    

    // 取票人证件类型
    
    FetchCertType   string `json:"fetch_cert_type,omitempty"`
    

    // 取票人电话
    
    FetchPhone   string `json:"fetch_phone,omitempty"`
    

    // 取票人姓名
    
    FetchTicketName   string `json:"fetch_ticket_name,omitempty"`
    

}
*/

// B2BFetchHolderInfo 
type B2BFetchHolderInfo struct {

    // 取票人证件号码
    FetchCertNumber   string `json:"fetch_cert_number,omitempty"`

    // 取票人证件类型
    FetchCertType   string `json:"fetch_cert_type,omitempty"`

    // 取票人电话
    FetchPhone   string `json:"fetch_phone,omitempty"`

    // 取票人姓名
    FetchTicketName   string `json:"fetch_ticket_name,omitempty"`

}
