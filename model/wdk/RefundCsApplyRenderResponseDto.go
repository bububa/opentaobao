package wdk

// RefundCsApplyRenderResponseDto 结构体
type RefundCsApplyRenderResponseDto struct {
	// 商家经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道主订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 请求唯一键, 提交请求(alibaba.tcls.aelophy.refund.csapply)时保持一致
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 申请退款的子订单列表
	OutSubOrders []Outsuborders `json:"out_sub_orders,omitempty" xml:"out_sub_orders>outsuborders,omitempty"`
	// 退款原因枚举列表
	ReasonList []Reasonlist `json:"reason_list,omitempty" xml:"reason_list>reasonlist,omitempty"`
}
