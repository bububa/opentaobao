package wdk

import (
	"sync"
)

// ReceiptQueryRequest 结构体
type ReceiptQueryRequest struct {
	// 打印批次
	BatchId string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 打印纸宽度
	PaperWidth int64 `json:"paper_width,omitempty" xml:"paper_width,omitempty"`
}

var poolReceiptQueryRequest = sync.Pool{
	New: func() any {
		return new(ReceiptQueryRequest)
	},
}

// GetReceiptQueryRequest() 从对象池中获取ReceiptQueryRequest
func GetReceiptQueryRequest() *ReceiptQueryRequest {
	return poolReceiptQueryRequest.Get().(*ReceiptQueryRequest)
}

// ReleaseReceiptQueryRequest 释放ReceiptQueryRequest
func ReleaseReceiptQueryRequest(v *ReceiptQueryRequest) {
	v.BatchId = ""
	v.PaperWidth = 0
	poolReceiptQueryRequest.Put(v)
}
