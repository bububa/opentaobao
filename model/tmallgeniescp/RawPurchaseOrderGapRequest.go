package tmallgeniescp

// RawPurchaseOrderGapRequest 结构体
type RawPurchaseOrderGapRequest struct {
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
	// 请求对象列表
	RawPurchaseOrderGapDTOs []RawPurchaseOrderGapDto `json:"raw_purchase_order_gap_d_t_os,omitempty" xml:"raw_purchase_order_gap_d_t_os>raw_purchase_order_gap_dto,omitempty"`
}
