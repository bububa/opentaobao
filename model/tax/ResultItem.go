package tax

import (
	"sync"
)

// ResultItem 结构体
type ResultItem struct {
	// 每一项具体异常信息
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
	// 每一项的请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 每一项成功失败
	Correctness bool `json:"correctness,omitempty" xml:"correctness,omitempty"`
}

var poolResultItem = sync.Pool{
	New: func() any {
		return new(ResultItem)
	},
}

// GetResultItem() 从对象池中获取ResultItem
func GetResultItem() *ResultItem {
	return poolResultItem.Get().(*ResultItem)
}

// ReleaseResultItem 释放ResultItem
func ReleaseResultItem(v *ResultItem) {
	v.ErrorDescription = ""
	v.RequestId = ""
	v.Correctness = false
	poolResultItem.Put(v)
}
