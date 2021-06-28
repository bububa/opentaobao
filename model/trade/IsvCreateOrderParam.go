package trade

// IsvCreateOrderParam 
/* model for simplify = false
type IsvCreateOrderParam struct {

    // 买家备注
    
    BuyerRemarks   string `json:"buyer_remarks,omitempty"`
    

    // 外部订单ID，仅仅做追踪使用
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 收货地址信息
    
    SmAddrModel  *struct {
        SmAddrModel  *SmAddrModel `json:"sm_addr_model,omitempty"`
    } `json:"sm_addr_model,omitempty"`
    

    // 订单行详情
    
    SubOrderInfoList  struct {
        IsvSimpleSubOrderModel  []IsvSimpleSubOrderModel `json:"isv_simple_sub_order_model,omitempty"`
    } `json:"sub_order_info_list,omitempty"`
    

}
*/

// IsvCreateOrderParam 
type IsvCreateOrderParam struct {

    // 买家备注
    BuyerRemarks   string `json:"buyer_remarks,omitempty"`

    // 外部订单ID，仅仅做追踪使用
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 收货地址信息
    SmAddrModel   *SmAddrModel `json:"sm_addr_model,omitempty"`

    // 订单行详情
    SubOrderInfoList   []IsvSimpleSubOrderModel `json:"sub_order_info_list,omitempty"`

}
