package tmallcar

// CreditLoanStatusSyncReq 结构体
type CreditLoanStatusSyncReq struct {
	// 金融方案id
	FinancePlanId string `json:"finance_plan_id,omitempty" xml:"finance_plan_id,omitempty"`
	// 网商订单id
	WsBankOrderId string `json:"ws_bank_order_id,omitempty" xml:"ws_bank_order_id,omitempty"`
	// 状态值
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 代扣参数
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
	// 交易场景码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 淘宝订单号id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 月供
	MonthlyPayment int64 `json:"monthly_payment,omitempty" xml:"monthly_payment,omitempty"`
	// 期数
	Periods int64 `json:"periods,omitempty" xml:"periods,omitempty"`
	// 首付金额
	DownPaymentAmount int64 `json:"down_payment_amount,omitempty" xml:"down_payment_amount,omitempty"`
	// 贷款金额
	CreditAmount int64 `json:"credit_amount,omitempty" xml:"credit_amount,omitempty"`
	// 放款金额
	LoanAmount int64 `json:"loan_amount,omitempty" xml:"loan_amount,omitempty"`
}
