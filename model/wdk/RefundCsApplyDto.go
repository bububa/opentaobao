package wdk

// RefundCsApplyDto 
/* model for simplify = false
type RefundCsApplyDto struct {

    // 渠道订单ID
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 商家经营店ID
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 请求唯一键
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 申请退款的子订单ID列表
    
    OutSubOrderIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"out_sub_order_ids,omitempty"`
    

    // 备注说明
    
    Memo   string `json:"memo,omitempty"`
    

    // 退款原因id
    
    ReasonId   int64 `json:"reason_id,omitempty"`
    

}
*/

// RefundCsApplyDto 
type RefundCsApplyDto struct {

    // 渠道订单ID
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 商家经营店ID
    StoreId   string `json:"store_id,omitempty"`

    // 请求唯一键
    RequestId   string `json:"request_id,omitempty"`

    // 申请退款的子订单ID列表
    OutSubOrderIds   []string `json:"out_sub_order_ids,omitempty"`

    // 备注说明
    Memo   string `json:"memo,omitempty"`

    // 退款原因id
    ReasonId   int64 `json:"reason_id,omitempty"`

}
