package cainiaohandover

// OpenFeeDetailDto 结构体
type OpenFeeDetailDto struct {
	// 已支付费用币种
	PaidFeeCurrency string `json:"paid_fee_currency,omitempty" xml:"paid_fee_currency,omitempty"`
	// 应支付费用币种
	FeeCurrency string `json:"fee_currency,omitempty" xml:"fee_currency,omitempty"`
	// 详细费用类型，normal_delivery_fee：配送费，sms_service_fee
	FeeDetailType string `json:"fee_detail_type,omitempty" xml:"fee_detail_type,omitempty"`
	// 已支付费用
	PaidFee int64 `json:"paid_fee,omitempty" xml:"paid_fee,omitempty"`
	// 应支付费用
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}
