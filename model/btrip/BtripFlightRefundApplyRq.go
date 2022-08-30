package btrip

// BtripFlightRefundApplyRq 结构体
type BtripFlightRefundApplyRq struct {
	// 乘客航段列表
	PassengerSegmentInfoList []PassengerSegmentInfo `json:"passenger_segment_info_list,omitempty" xml:"passenger_segment_info_list>passenger_segment_info,omitempty"`
	// 退票凭证
	RefundVoucherInfo []string `json:"refund_voucher_info,omitempty" xml:"refund_voucher_info>string,omitempty"`
	// 退票票号列表
	TicketNos []string `json:"ticket_nos,omitempty" xml:"ticket_nos>string,omitempty"`
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 退票理由
	ReasonDetail string `json:"reason_detail,omitempty" xml:"reason_detail,omitempty"`
	// 分销子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 分销外部退票申请单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 商品单元
	ItemUnitIds string `json:"item_unit_ids,omitempty" xml:"item_unit_ids,omitempty"`
	// 会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 订单退款总金额(单位元)
	DisplayRefundMoney string `json:"display_refund_money,omitempty" xml:"display_refund_money,omitempty"`
	// 是否自愿
	IsVoluntary int64 `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// 退票原因类型
	ReasonType int64 `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
	// 退票原因类型
	ReasoType int64 `json:"reaso_type,omitempty" xml:"reaso_type,omitempty"`
	// 订单退款总金额
	TotalRefundPrice int64 `json:"total_refund_price,omitempty" xml:"total_refund_price,omitempty"`
	// 企业退款金额
	CorpRefundPrice int64 `json:"corp_refund_price,omitempty" xml:"corp_refund_price,omitempty"`
	// 个人退款金额
	PersonalRefundPrice int64 `json:"personal_refund_price,omitempty" xml:"personal_refund_price,omitempty"`
}
