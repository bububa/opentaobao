package bus

// AgentConfirmBookPassengerInfo 
/* model for simplify = false
type AgentConfirmBookPassengerInfo struct {

    // 票ID
    
    AgentTicketId   string `json:"agent_ticket_id,omitempty"`
    

    // 证件号码
    
    PassengerCertNo   string `json:"passenger_cert_no,omitempty"`
    

    // 证件类型:01-身份证;02-护照;03-港澳通行证;04-台湾通行证
    
    PassengerCertType   string `json:"passenger_cert_type,omitempty"`
    

    // 乘客名称
    
    PassengerName   string `json:"passenger_name,omitempty"`
    

    // 座位号
    
    SeatNo   string `json:"seat_no,omitempty"`
    

}
*/

// AgentConfirmBookPassengerInfo 
type AgentConfirmBookPassengerInfo struct {

    // 票ID
    AgentTicketId   string `json:"agent_ticket_id,omitempty"`

    // 证件号码
    PassengerCertNo   string `json:"passenger_cert_no,omitempty"`

    // 证件类型:01-身份证;02-护照;03-港澳通行证;04-台湾通行证
    PassengerCertType   string `json:"passenger_cert_type,omitempty"`

    // 乘客名称
    PassengerName   string `json:"passenger_name,omitempty"`

    // 座位号
    SeatNo   string `json:"seat_no,omitempty"`

}
