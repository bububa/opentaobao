package tbtrade

// IdentifyLogisticsInfo 结构体
type IdentifyLogisticsInfo struct {
	// 物流公司
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 运单号
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 退款单号
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 子订单号
	DetailOrderId int64 `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
	// 阶段号
	StageNo int64 `json:"stage_no,omitempty" xml:"stage_no,omitempty"`
	// 是否退货
	Refund bool `json:"refund,omitempty" xml:"refund,omitempty"`
}
