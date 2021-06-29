package wdk

// RefundCsApplyDto 
type RefundCsApplyDto struct {
    // 渠道订单ID
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 商家经营店ID
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 请求唯一键
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 申请退款的子订单ID列表
    OutSubOrderIds   []string `json:"out_sub_order_ids,omitempty" xml:"out_sub_order_ids>string,omitempty"`
    // 备注说明
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    // 退款原因id
    ReasonId   int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}
