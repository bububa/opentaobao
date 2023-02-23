package alitripmerchant

// DerbyVoucherReceiptApplyVo 结构体
type DerbyVoucherReceiptApplyVo struct {
	// 发票申请流水号
	FlowNumber string `json:"flow_number,omitempty" xml:"flow_number,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
