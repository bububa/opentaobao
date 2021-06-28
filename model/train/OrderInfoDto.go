package train

// OrderInfoDto 
type OrderInfoDto struct {

    // 乘车人信息
    
    PassengerList   []TrainPassengerInfoDto `json:"passenger_list,omitempty" xml:"passenger_list,omitempty"`
    

    // 是否联程
    
    IsMultiTrip   bool `json:"is_multi_trip,omitempty" xml:"is_multi_trip,omitempty"`
    

    // 是否可以退改
    
    CanRefund   bool `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
    

    // 12306支付截止时间
    
    LimitPayTime   string `json:"limit_pay_time,omitempty" xml:"limit_pay_time,omitempty"`
    

    // 票号
    
    TicketNo   string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
    

    // 失败子订单号
    
    FailSubOrderId   int64 `json:"fail_sub_order_id,omitempty" xml:"fail_sub_order_id,omitempty"`
    

    // 主订单号
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 支付url
    
    PayUrl   string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
    

    // 车次信息
    
    BaseDo   *TrainBaseDto `json:"base_do,omitempty" xml:"base_do,omitempty"`
    

}
