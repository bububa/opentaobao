package ascpchannel

// Requestparams 
type Requestparams struct {

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 外部订单号
    
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 金额
    
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    

}
