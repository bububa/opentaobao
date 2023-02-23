package flightuppc

// InsOrderOpenPersonDto 结构体
type InsOrderOpenPersonDto struct {
	// idCardNo(md5脱敏)
	IdCardNo string `json:"id_card_no,omitempty" xml:"id_card_no,omitempty"`
	// 子保单号
	PolicyNo string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 保险订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// idCardType
	IdCardType int64 `json:"id_card_type,omitempty" xml:"id_card_type,omitempty"`
	// 外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}
