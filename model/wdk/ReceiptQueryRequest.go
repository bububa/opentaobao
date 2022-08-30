package wdk

// ReceiptQueryRequest 结构体
type ReceiptQueryRequest struct {
	// 打印批次
	BatchId string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 打印纸宽度
	PaperWidth int64 `json:"paper_width,omitempty" xml:"paper_width,omitempty"`
}
