package alisports

// SubsidyAmountVo 结构体
type SubsidyAmountVo struct {
	// 用户的支付宝ID
	AlipayId string `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
	// 补助金额
	SubsidyAmount int64 `json:"subsidy_amount,omitempty" xml:"subsidy_amount,omitempty"`
}
