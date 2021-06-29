package omniorder

// TradeOrderDetailDto 
type TradeOrderDetailDto struct {

    // 商品信息，非淘及散件必填，淘系订单以订单在淘宝系统内容为准
    
    Items   []TradeItemInfoDto `json:"items,omitempty" xml:"items,omitempty"`
    

    // 子订单ID，如果是单一订单，则主订单ID和子订单ID均填入同一值
    
    SubOrderId   string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    

}
