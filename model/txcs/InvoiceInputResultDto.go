package txcs

// InvoiceInputResultDto 
type InvoiceInputResultDto struct {
    // 失败原因
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 发票代码
    InvoiceCode   string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
    // 发票号码
    InvoiceNo   string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
}
