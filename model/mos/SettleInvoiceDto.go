package mos

// SettleInvoiceDto 结构体
type SettleInvoiceDto struct {
	// 发票类型 SPECIAL("专票")，ORDINARY("普票")
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 发票行号
	InvoicelineNo string `json:"invoiceline_no,omitempty" xml:"invoiceline_no,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 发票不含税金额
	ExcludingTaxAmount string `json:"excluding_tax_amount,omitempty" xml:"excluding_tax_amount,omitempty"`
	// 结算行税率，普票的税率必须是0
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 税额，普票时填0
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 发票总金额， 必须=excludingTaxAmount+taxAmount
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 开票日期
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// 扩展
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}
