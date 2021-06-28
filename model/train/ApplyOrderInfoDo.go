package train

// ApplyOrderInfoDo 
/* model for simplify = false
type ApplyOrderInfoDo struct {

    // 退款金额（单位分）
    
    RefundFee   int64 `json:"refund_fee,omitempty"`
    

    // 订单号
    
    OrderId   int64 `json:"order_id,omitempty"`
    

    // 乘车人信息
    
    PassengerList  struct {
        TrainPassengerInfoDto  []TrainPassengerInfoDto `json:"train_passenger_info_dto,omitempty"`
    } `json:"passenger_list,omitempty"`
    

    // 支付url
    
    PayUrl   string `json:"pay_url,omitempty"`
    

    // 是否可以退改 true:支持 false:不支持
    
    CanRefund   bool `json:"can_refund,omitempty"`
    

    // 改签单号
    
    ApplyId   int64 `json:"apply_id,omitempty"`
    

    // 改签支付url有效期
    
    LimitAlipayTime   string `json:"limit_alipay_time,omitempty"`
    

    // 取票号
    
    TicketNo   string `json:"ticket_no,omitempty"`
    

    // 改签手续费（单位分）
    
    HandFee   int64 `json:"hand_fee,omitempty"`
    

    // 支付超时时间
    
    LimitPayTime   string `json:"limit_pay_time,omitempty"`
    

}
*/

// ApplyOrderInfoDo 
type ApplyOrderInfoDo struct {

    // 退款金额（单位分）
    RefundFee   int64 `json:"refund_fee,omitempty"`

    // 订单号
    OrderId   int64 `json:"order_id,omitempty"`

    // 乘车人信息
    PassengerList   []TrainPassengerInfoDto `json:"passenger_list,omitempty"`

    // 支付url
    PayUrl   string `json:"pay_url,omitempty"`

    // 是否可以退改 true:支持 false:不支持
    CanRefund   bool `json:"can_refund,omitempty"`

    // 改签单号
    ApplyId   int64 `json:"apply_id,omitempty"`

    // 改签支付url有效期
    LimitAlipayTime   string `json:"limit_alipay_time,omitempty"`

    // 取票号
    TicketNo   string `json:"ticket_no,omitempty"`

    // 改签手续费（单位分）
    HandFee   int64 `json:"hand_fee,omitempty"`

    // 支付超时时间
    LimitPayTime   string `json:"limit_pay_time,omitempty"`

}
