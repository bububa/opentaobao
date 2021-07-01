package tbk

// TljInstanceReportDto 结构体
type TljInstanceReportDto struct {
	// 解冻金额
	UnfreezeAmount string `json:"unfreeze_amount,omitempty" xml:"unfreeze_amount,omitempty"`
	// 解冻红包个数
	UnfreezeNum int64 `json:"unfreeze_num,omitempty" xml:"unfreeze_num,omitempty"`
	// 失效回退金额
	RefundAmount string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 失效回退红包个数
	RefundNum int64 `json:"refund_num,omitempty" xml:"refund_num,omitempty"`
	// 引导成交金额
	AlipayAmount string `json:"alipay_amount,omitempty" xml:"alipay_amount,omitempty"`
	// 红包核销金额
	UseAmount string `json:"use_amount,omitempty" xml:"use_amount,omitempty"`
	// 红包核销个数
	UseNum int64 `json:"use_num,omitempty" xml:"use_num,omitempty"`
	// 红包领取金额
	WinAmount string `json:"win_amount,omitempty" xml:"win_amount,omitempty"`
	// 红包领取个数
	WinNum int64 `json:"win_num,omitempty" xml:"win_num,omitempty"`
	// 引导预估佣金金额
	PreCommissionAmount string `json:"pre_commission_amount,omitempty" xml:"pre_commission_amount,omitempty"`
	// 退款红包金额
	FpRefundAmount string `json:"fp_refund_amount,omitempty" xml:"fp_refund_amount,omitempty"`
	// 退款红包个数
	FpRefundNum int64 `json:"fp_refund_num,omitempty" xml:"fp_refund_num,omitempty"`
}
