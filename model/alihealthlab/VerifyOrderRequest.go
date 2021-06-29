package alihealthlab

// VerifyOrderRequest 
type VerifyOrderRequest struct {

    // isv的预约单 id
    
    IsvReserveId   string `json:"isv_reserve_id,omitempty" xml:"isv_reserve_id,omitempty"`
    

    // 需要核销的订单 id
    
    BizOrderIds   []string `json:"biz_order_ids,omitempty" xml:"biz_order_ids>string,omitempty"`
    

    // 健康侧预约单号
    
    JkReserveId   string `json:"jk_reserve_id,omitempty" xml:"jk_reserve_id,omitempty"`
    

}
