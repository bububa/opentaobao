package legalsuit

// SuitFeeModel 结构体
type SuitFeeModel struct {
	// 财产保全保险费
	DepositInsuranceAmount string `json:"deposit_insurance_amount,omitempty" xml:"deposit_insurance_amount,omitempty"`
	// 其他费用
	OtherFeeAmount string `json:"other_fee_amount,omitempty" xml:"other_fee_amount,omitempty"`
	// 评估鉴定费
	EvaluationFeeAmount string `json:"evaluation_fee_amount,omitempty" xml:"evaluation_fee_amount,omitempty"`
	// 公证费
	NotarialFeeAmount string `json:"notarial_fee_amount,omitempty" xml:"notarial_fee_amount,omitempty"`
	// 公告费
	AnnouncementFeeAmount string `json:"announcement_fee_amount,omitempty" xml:"announcement_fee_amount,omitempty"`
	// 保证金金额
	SecurityDepositAmount string `json:"security_deposit_amount,omitempty" xml:"security_deposit_amount,omitempty"`
	// 退还诉讼费金额
	ReturnCourtFeeAmount string `json:"return_court_fee_amount,omitempty" xml:"return_court_fee_amount,omitempty"`
	// 保全费金额
	PreserveFeeAmount string `json:"preserve_fee_amount,omitempty" xml:"preserve_fee_amount,omitempty"`
	// 补缴诉讼费金额
	MakeupPayCourtFeeAmount string `json:"makeup_pay_court_fee_amount,omitempty" xml:"makeup_pay_court_fee_amount,omitempty"`
	// 补缴诉讼费日期
	MakeupPayCourtFeeTime string `json:"makeup_pay_court_fee_time,omitempty" xml:"makeup_pay_court_fee_time,omitempty"`
	// 预缴诉讼费金额
	PrepayCourtFeeAmount string `json:"prepay_court_fee_amount,omitempty" xml:"prepay_court_fee_amount,omitempty"`
	// 预缴诉讼费日期
	PrepayCourtFeeTime string `json:"prepay_court_fee_time,omitempty" xml:"prepay_court_fee_time,omitempty"`
}
