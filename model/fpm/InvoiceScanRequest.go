package fpm

// InvoiceScanRequest 结构体
type InvoiceScanRequest struct {
	// 发票实体
	InvoiceScanShareData []XforceInvoiceDto `json:"invoice_scan_share_data,omitempty" xml:"invoice_scan_share_data>xforce_invoice_dto,omitempty"`
	// 签名值
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 应用code
	Appcode string `json:"appcode,omitempty" xml:"appcode,omitempty"`
	// 扫描时间
	T int64 `json:"_t,omitempty" xml:"_t,omitempty"`
}
