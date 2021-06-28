package tbk

// OrderData 
type OrderData struct {

    // 预估佣金
    
    Commission   string `json:"commission,omitempty" xml:"commission,omitempty"`
    

    // 收货时间
    
    ConfirmReceiveTime   string `json:"confirm_receive_time,omitempty" xml:"confirm_receive_time,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 订单号
    
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    

}
