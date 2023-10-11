package ascp

// InquiryResult 结构体
type InquiryResult struct {
	// BFC单号
	WdsCoordinationOrderId string `json:"wds_coordination_order_id,omitempty" xml:"wds_coordination_order_id,omitempty"`
	// 服务商id
	TpId string `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 询价结果明细
	WorkerPriceInfo *WorkerPriceInfo `json:"worker_price_info,omitempty" xml:"worker_price_info,omitempty"`
}
