package idleisv

import (
	"sync"
)

// AlibabaIdleIsvItemDownshelfTopResult 结构体
type AlibabaIdleIsvItemDownshelfTopResult struct {
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaIdleIsvItemDownshelfTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemDownshelfTopResult)
	},
}

// GetAlibabaIdleIsvItemDownshelfTopResult() 从对象池中获取AlibabaIdleIsvItemDownshelfTopResult
func GetAlibabaIdleIsvItemDownshelfTopResult() *AlibabaIdleIsvItemDownshelfTopResult {
	return poolAlibabaIdleIsvItemDownshelfTopResult.Get().(*AlibabaIdleIsvItemDownshelfTopResult)
}

// ReleaseAlibabaIdleIsvItemDownshelfTopResult 释放AlibabaIdleIsvItemDownshelfTopResult
func ReleaseAlibabaIdleIsvItemDownshelfTopResult(v *AlibabaIdleIsvItemDownshelfTopResult) {
	v.Data = false
	poolAlibabaIdleIsvItemDownshelfTopResult.Put(v)
}
