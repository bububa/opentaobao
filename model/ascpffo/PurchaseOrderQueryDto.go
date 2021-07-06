package ascpffo

// PurchaseOrderQueryDto 结构体
type PurchaseOrderQueryDto struct {
	// 采购单号
	PurchaseOrderNo string `json:"purchase_order_no,omitempty" xml:"purchase_order_no,omitempty"`
	// 行业账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}
