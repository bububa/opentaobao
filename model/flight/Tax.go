package flight

// Tax 结构体
type Tax struct {
	// 税费编码(CN:机场建设费 ,YQ	燃油附加税)
	TaxCode string `json:"tax_code,omitempty" xml:"tax_code,omitempty"`
	// 税费金额，单位:分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 不可退税费，单位:分
	NotRefundAmount int64 `json:"not_refund_amount,omitempty" xml:"not_refund_amount,omitempty"`
}
