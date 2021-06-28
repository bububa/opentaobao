package train

// OrderTicketInfo 
/* model for simplify = false
type OrderTicketInfo struct {

    // ttp子单ID
    
    TtpSubOrderId   int64 `json:"ttp_sub_order_id,omitempty"`
    

    // 真实票价
    
    RealTicketPrice   int64 `json:"real_ticket_price,omitempty"`
    

    // 真实坐席
    
    RealSeat   int64 `json:"real_seat,omitempty"`
    

    // 座位号
    
    SeatNum   string `json:"seat_num,omitempty"`
    

    // 车次
    
    TrainNo   string `json:"train_no,omitempty"`
    

    // 乘客姓名
    
    PassengerName   string `json:"passenger_name,omitempty"`
    

    // 证件类型
    
    CertType   string `json:"cert_type,omitempty"`
    

    // 证件号
    
    CertificateNum   string `json:"certificate_num,omitempty"`
    

    // 保险支付金额
    
    InsurancePayPrice   int64 `json:"insurance_pay_price,omitempty"`
    

    // 定制票出票结果 1:定制票出票 0:非定制票出票
    
    VipCustomResult   int64 `json:"vip_custom_result,omitempty"`
    

    // 12306票号
    
    TicketNo   string `json:"ticket_no,omitempty"`
    

}
*/

// OrderTicketInfo 
type OrderTicketInfo struct {

    // ttp子单ID
    TtpSubOrderId   int64 `json:"ttp_sub_order_id,omitempty"`

    // 真实票价
    RealTicketPrice   int64 `json:"real_ticket_price,omitempty"`

    // 真实坐席
    RealSeat   int64 `json:"real_seat,omitempty"`

    // 座位号
    SeatNum   string `json:"seat_num,omitempty"`

    // 车次
    TrainNo   string `json:"train_no,omitempty"`

    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 证件类型
    CertType   string `json:"cert_type,omitempty"`

    // 证件号
    CertificateNum   string `json:"certificate_num,omitempty"`

    // 保险支付金额
    InsurancePayPrice   int64 `json:"insurance_pay_price,omitempty"`

    // 定制票出票结果 1:定制票出票 0:非定制票出票
    VipCustomResult   int64 `json:"vip_custom_result,omitempty"`

    // 12306票号
    TicketNo   string `json:"ticket_no,omitempty"`

}
