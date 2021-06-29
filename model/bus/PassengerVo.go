package bus

// PassengerVo 
type PassengerVo struct {
    // 乘客证件号码
    RiderCertNumber   string `json:"rider_cert_number,omitempty" xml:"rider_cert_number,omitempty"`
    // 乘客证件类型
    RiderCertType   string `json:"rider_cert_type,omitempty" xml:"rider_cert_type,omitempty"`
    // 乘客姓名
    RiderName   string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
    // 服务费
    ServiceCharge   int64 `json:"service_charge,omitempty" xml:"service_charge,omitempty"`
    // 票价
    TicketPrice   int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}
