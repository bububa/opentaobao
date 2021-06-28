package bus

// B2BTicketInfo 
/* model for simplify = false
type B2BTicketInfo struct {

    // 乘客类型
    
    RiderCertType   string `json:"rider_cert_type,omitempty"`
    

    // 乘客姓名
    
    RiderName   string `json:"rider_name,omitempty"`
    

    // 座位号
    
    RiderSeatNumber   string `json:"rider_seat_number,omitempty"`
    

    // 服务费
    
    ServiceCharge   int64 `json:"service_charge,omitempty"`
    

    // 票ID
    
    TicketId   string `json:"ticket_id,omitempty"`
    

    // 票价
    
    TicketPrice   int64 `json:"ticket_price,omitempty"`
    

    // 退票手续费
    
    CommissionFee   int64 `json:"commission_fee,omitempty"`
    

    // 可退款金额
    
    RefundFee   int64 `json:"refund_fee,omitempty"`
    

    // 退票状态 1(新建立) 2（处理中）3（同意）4（拒绝）5（已退款）
    
    RefundStatus   int64 `json:"refund_status,omitempty"`
    

    // 子订单id
    
    SubOrderId   int64 `json:"sub_order_id,omitempty"`
    

}
*/

// B2BTicketInfo 
type B2BTicketInfo struct {

    // 乘客类型
    RiderCertType   string `json:"rider_cert_type,omitempty"`

    // 乘客姓名
    RiderName   string `json:"rider_name,omitempty"`

    // 座位号
    RiderSeatNumber   string `json:"rider_seat_number,omitempty"`

    // 服务费
    ServiceCharge   int64 `json:"service_charge,omitempty"`

    // 票ID
    TicketId   string `json:"ticket_id,omitempty"`

    // 票价
    TicketPrice   int64 `json:"ticket_price,omitempty"`

    // 退票手续费
    CommissionFee   int64 `json:"commission_fee,omitempty"`

    // 可退款金额
    RefundFee   int64 `json:"refund_fee,omitempty"`

    // 退票状态 1(新建立) 2（处理中）3（同意）4（拒绝）5（已退款）
    RefundStatus   int64 `json:"refund_status,omitempty"`

    // 子订单id
    SubOrderId   int64 `json:"sub_order_id,omitempty"`

}
