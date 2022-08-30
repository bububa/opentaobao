package mos

// PosPaymentDto 结构体
type PosPaymentDto struct {
	// 扩展参数。注：POS中的支付大类、小类，通过extendParams传。支付大类属性名： paymentType，支付小类属性名： paymentSubType。
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 支付代码
	PaymentCode string `json:"payment_code,omitempty" xml:"payment_code,omitempty"`
	// 支付名称
	PaymentName string `json:"payment_name,omitempty" xml:"payment_name,omitempty"`
	// 支付流水号
	PaymentNo string `json:"payment_no,omitempty" xml:"payment_no,omitempty"`
	// 支付/优惠 金额,单位：分
	PaymentAmount int64 `json:"payment_amount,omitempty" xml:"payment_amount,omitempty"`
	// 支付行号
	PaymentLineNo int64 `json:"payment_line_no,omitempty" xml:"payment_line_no,omitempty"`
	// 1:支付，2:优惠
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
}
