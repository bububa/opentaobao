package fpm

// XforceInvoiceDto 结构体
type XforceInvoiceDto struct {
	// 发票行
	DetailList []InvoiceDetails `json:"detail_list,omitempty" xml:"detail_list>invoice_details,omitempty"`
	// 发票头
	Head *InvoiceMainExt `json:"head,omitempty" xml:"head,omitempty"`
}
