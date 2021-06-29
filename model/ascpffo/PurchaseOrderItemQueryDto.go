package ascpffo

// PurchaseOrderItemQueryDto 
type PurchaseOrderItemQueryDto struct {
    // 账套编码
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 分页页码
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    // 分页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 采购单号
    PurchaseOrderNo   string `json:"purchase_order_no,omitempty" xml:"purchase_order_no,omitempty"`
}
