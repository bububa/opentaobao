package einvoice

// InvoiceItems 
type InvoiceItems struct {
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 数量
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 价税合计
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 规格型号
    Specification   string `json:"specification,omitempty" xml:"specification,omitempty"`
    // 单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
}
