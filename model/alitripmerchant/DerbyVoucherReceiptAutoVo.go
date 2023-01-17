package alitripmerchant

// DerbyVoucherReceiptAutoVo 结构体
type DerbyVoucherReceiptAutoVo struct {
	// 详情
	Data []DerbyVoucherReceiptAutoResponse `json:"data,omitempty" xml:"data>derby_voucher_receipt_auto_response,omitempty"`
	// 发票抬头
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 税号
	TaxID string `json:"tax_i_d,omitempty" xml:"tax_i_d,omitempty"`
}
