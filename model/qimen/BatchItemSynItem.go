package qimen

import (
	"sync"
)

// BatchItemSynItem 结构体
type BatchItemSynItem struct {
	// 没有同步成功的商品的编码
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 出错信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolBatchItemSynItem = sync.Pool{
	New: func() any {
		return new(BatchItemSynItem)
	},
}

// GetBatchItemSynItem() 从对象池中获取BatchItemSynItem
func GetBatchItemSynItem() *BatchItemSynItem {
	return poolBatchItemSynItem.Get().(*BatchItemSynItem)
}

// ReleaseBatchItemSynItem 释放BatchItemSynItem
func ReleaseBatchItemSynItem(v *BatchItemSynItem) {
	v.ItemCode = ""
	v.Message = ""
	poolBatchItemSynItem.Put(v)
}
