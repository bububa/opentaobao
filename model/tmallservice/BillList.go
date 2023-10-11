package tmallservice

// BillList 结构体
type BillList struct {
	// 工单费用清单
	FeeList []FeeList `json:"fee_list,omitempty" xml:"fee_list>fee_list,omitempty"`
	// 费用金额
	FeeAmount string `json:"fee_amount,omitempty" xml:"fee_amount,omitempty"`
	// 费用备注
	FeeNotice string `json:"fee_notice,omitempty" xml:"fee_notice,omitempty"`
	// 费用来源单号，仅退款和增加费用有值
	SrcOrderId string `json:"src_order_id,omitempty" xml:"src_order_id,omitempty"`
	// 提现时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 抽佣比例
	PlatformCommissionRate string `json:"platform_commission_rate,omitempty" xml:"platform_commission_rate,omitempty"`
	// 费用名称
	FeeName string `json:"fee_name,omitempty" xml:"fee_name,omitempty"`
	// 提现支付宝流水号
	PayTradeNo string `json:"pay_trade_no,omitempty" xml:"pay_trade_no,omitempty"`
	// 费用产生时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 费用类型
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 提现支付宝备注
	PayTradeNotice string `json:"pay_trade_notice,omitempty" xml:"pay_trade_notice,omitempty"`
	// 费用产生时间
	BillTime string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
	// 总退款金额
	SumRefundAmount *BigDecimal `json:"sum_refund_amount,omitempty" xml:"sum_refund_amount,omitempty"`
	// 总增加费用金额
	SumAddAmount *BigDecimal `json:"sum_add_amount,omitempty" xml:"sum_add_amount,omitempty"`
	// 提现金额
	PayAmount *BigDecimal `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 分成金额
	FcAmount *BigDecimal `json:"fc_amount,omitempty" xml:"fc_amount,omitempty"`
	// 总服务费用金额
	SumServiceAmount *BigDecimal `json:"sum_service_amount,omitempty" xml:"sum_service_amount,omitempty"`
}
