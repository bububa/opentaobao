package wdk

// RefundCsApplyRenderDTO 
type RefundCsApplyRenderDTO struct {
    // 渠道订单ID
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 商家经营店ID
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 申请退款的子订单ID列表
    OutSubOrderIds   []string `json:"out_sub_order_ids,omitempty" xml:"out_sub_order_ids>string,omitempty"`
}
