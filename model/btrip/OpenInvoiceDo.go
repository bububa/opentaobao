package btrip

// OpenInvoiceDO 
type OpenInvoiceDO struct {
    // 商旅发票id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 发票抬头
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 第三方发票id
    ThirdPartInvoiceId   string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
    // 发票类型：1:增值税普通发票；2:增值税专用发票
    InvoiceType   int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
}
