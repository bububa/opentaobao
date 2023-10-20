package scbp

import (
	"sync"
)

// ForbiddenKeywordOperation 结构体
type ForbiddenKeywordOperation struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

var poolForbiddenKeywordOperation = sync.Pool{
	New: func() any {
		return new(ForbiddenKeywordOperation)
	},
}

// GetForbiddenKeywordOperation() 从对象池中获取ForbiddenKeywordOperation
func GetForbiddenKeywordOperation() *ForbiddenKeywordOperation {
	return poolForbiddenKeywordOperation.Get().(*ForbiddenKeywordOperation)
}

// ReleaseForbiddenKeywordOperation 释放ForbiddenKeywordOperation
func ReleaseForbiddenKeywordOperation(v *ForbiddenKeywordOperation) {
	v.Keyword = ""
	poolForbiddenKeywordOperation.Put(v)
}
