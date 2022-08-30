package wdk

// ReceiptDto 结构体
type ReceiptDto struct {
	// 商家/顾客联小票数据
	ReceiptInfo *ReceiptInfoDto `json:"receipt_info,omitempty" xml:"receipt_info,omitempty"`
}
