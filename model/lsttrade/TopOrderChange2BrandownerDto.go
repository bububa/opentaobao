package lsttrade

// TopOrderChange2BrandownerDto 
type TopOrderChange2BrandownerDto struct {

    // 退款单id
    
    RefundId   int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    

    // 是否新建，包括订单新建和退款单新建
    
    NewCreated   bool `json:"new_created,omitempty" xml:"new_created,omitempty"`
    

    // FORWARD_ORDER 正向订单表示正常购买流程,REVERSE_ORDER 逆向订单表示退款流程
    
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    

    // 主订单id
    
    MainOrderId   int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    

    // 子订单id
    
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    

}
