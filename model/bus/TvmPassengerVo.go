package bus

// TvmPassengerVo 
type TvmPassengerVo struct {
    // 电子票内容
    AgentEticket   string `json:"agent_eticket,omitempty" xml:"agent_eticket,omitempty"`
    // 票号ID
    AgentTicketId   string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
    // 标准票价
    FullPrice   int64 `json:"full_price,omitempty" xml:"full_price,omitempty"`
    // 是否带有儿童
    HasChildren   bool `json:"has_children,omitempty" xml:"has_children,omitempty"`
    // 身份证号
    RiderCertNumber   string `json:"rider_cert_number,omitempty" xml:"rider_cert_number,omitempty"`
    // 证件类型
    RiderCertType   string `json:"rider_cert_type,omitempty" xml:"rider_cert_type,omitempty"`
    // 姓名
    RiderName   string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
    // 座位号
    SeatNumber   string `json:"seat_number,omitempty" xml:"seat_number,omitempty"`
    // 服务费
    ServiceCharge   int64 `json:"service_charge,omitempty" xml:"service_charge,omitempty"`
    // 实际票价
    TicketPrice   int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
    // 保险费 (单位分)
    InsurePrice   int64 `json:"insure_price,omitempty" xml:"insure_price,omitempty"`
    // 保险信息
    TvmInsuranceInfos   []TvmInsuranceInfo `json:"tvm_insurance_infos,omitempty" xml:"tvm_insurance_infos>tvm_insurance_info,omitempty"`
}
