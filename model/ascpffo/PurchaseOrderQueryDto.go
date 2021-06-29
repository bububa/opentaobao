package ascpffo

// PurchaseOrderQueryDTO 
type PurchaseOrderQueryDTO struct {
    // 行业账套编码
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 采购单号
    PurchaseOrderNo   string `json:"purchase_order_no,omitempty" xml:"purchase_order_no,omitempty"`
}
