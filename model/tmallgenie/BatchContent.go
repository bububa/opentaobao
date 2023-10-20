package tmallgenie

import (
	"sync"
)

// BatchContent 结构体
type BatchContent struct {
	// 内容信息
	OpenContents []OpenContent `json:"open_contents,omitempty" xml:"open_contents>open_content,omitempty"`
	// 类目ID，具体参见开放平台类目相关描述
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

var poolBatchContent = sync.Pool{
	New: func() any {
		return new(BatchContent)
	},
}

// GetBatchContent() 从对象池中获取BatchContent
func GetBatchContent() *BatchContent {
	return poolBatchContent.Get().(*BatchContent)
}

// ReleaseBatchContent 释放BatchContent
func ReleaseBatchContent(v *BatchContent) {
	v.OpenContents = v.OpenContents[:0]
	v.CategoryId = 0
	poolBatchContent.Put(v)
}
