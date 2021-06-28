package servicecenter

// TpFundsSendQuery 
type TpFundsSendQuery struct {

    // 资金权益类型。1：预付款红包；2：尾款红包
    
    FundsType   int64 `json:"funds_type,omitempty" xml:"funds_type,omitempty"`
    

    // 订单ID
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

}
