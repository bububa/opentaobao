package moscm

// InvoiceDto 结构体
type InvoiceDto struct {
	// 发票类型:普通发票、电子发票、增值税发票
	Kind string `json:"kind,omitempty" xml:"kind,omitempty"`
	// 名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 发票内容:日用品
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 发票抬头类型：个人、单位
	TitleType string `json:"title_type,omitempty" xml:"title_type,omitempty"`
	// 税号
	TaxNumber string `json:"tax_number,omitempty" xml:"tax_number,omitempty"`
}
