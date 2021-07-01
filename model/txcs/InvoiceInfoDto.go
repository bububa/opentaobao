package txcs

// InvoiceInfoDto 结构体
type InvoiceInfoDto struct {
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
}
