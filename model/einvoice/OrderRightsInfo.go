package einvoice

// OrderRightsInfo 结构体
type OrderRightsInfo struct {
	// 订单号
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 订单应开票时间
	ExceptInvoiceTime string `json:"except_invoice_time,omitempty" xml:"except_invoice_time,omitempty"`
}
