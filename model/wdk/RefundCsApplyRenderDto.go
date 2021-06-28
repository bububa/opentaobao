package wdk

// RefundCsApplyRenderDto 
/* model for simplify = false
type RefundCsApplyRenderDto struct {

    // 渠道订单ID
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 商家经营店ID
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 申请退款的子订单ID列表
    
    OutSubOrderIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"out_sub_order_ids,omitempty"`
    

}
*/

// RefundCsApplyRenderDto 
type RefundCsApplyRenderDto struct {

    // 渠道订单ID
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 商家经营店ID
    StoreId   string `json:"store_id,omitempty"`

    // 申请退款的子订单ID列表
    OutSubOrderIds   []string `json:"out_sub_order_ids,omitempty"`

}
