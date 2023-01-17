package bus

// MerchantBusRefundApplyInfo 结构体
type MerchantBusRefundApplyInfo struct {
	// 退款时间
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// 退票原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 扩展属性
	ExtAttr string `json:"ext_attr,omitempty" xml:"ext_attr,omitempty"`
	// 退票申请创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 同意/拒绝时间
	AgreeOrRefuseTime string `json:"agree_or_refuse_time,omitempty" xml:"agree_or_refuse_time,omitempty"`
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 拒绝原因
	RefuseReason string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
	// 0-用户申请1-商家申请2-二次退款
	ApplyType int64 `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 退票款金额，单位分
	RefundTicketAmount int64 `json:"refund_ticket_amount,omitempty" xml:"refund_ticket_amount,omitempty"`
	// 退款总金额，单位分
	RefundTotalAmount int64 `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty"`
	// 退票状态：1：新申请 2：代理商处理中 3：同意 4：拒绝 5:已退款
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退票申请id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 票信息
	BusTicketInfo *MerchantBusTicketInfo `json:"bus_ticket_info,omitempty" xml:"bus_ticket_info,omitempty"`
	// 退票手续费，单位分
	CommissionChargeAmount int64 `json:"commission_charge_amount,omitempty" xml:"commission_charge_amount,omitempty"`
	// 退商家服务费金额，单位分
	RefundServiceChargeAmount int64 `json:"refund_service_charge_amount,omitempty" xml:"refund_service_charge_amount,omitempty"`
}
