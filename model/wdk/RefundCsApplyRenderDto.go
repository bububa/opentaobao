package wdk

// RefundCsApplyRenderDto 
type RefundCsApplyRenderDto struct {

    // 渠道订单ID
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 商家经营店ID
    StoreId   string `json:"store_id,omitempty"`

    // 申请退款的子订单ID列表
    OutSubOrderIds   []String `json:"out_sub_order_ids,omitempty"`

}
