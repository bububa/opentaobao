package westcrm

// UserConsumerVo 
type UserConsumerVo struct {

    // 订单
    
    Orders   []OrderVo `json:"orders,omitempty" xml:"orders,omitempty"`
    

    // 消费总额
    
    TotalPay   string `json:"total_pay,omitempty" xml:"total_pay,omitempty"`
    

    // 用户id
    
    IbUserId   int64 `json:"ib_user_id,omitempty" xml:"ib_user_id,omitempty"`
    

}
