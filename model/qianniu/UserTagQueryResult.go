package qianniu

import (
	"sync"
)

// UserTagQueryResult 结构体
type UserTagQueryResult struct {
	// 买家是否有这个标，true表示有，false表示没有
	EnterpriseBuyer bool `json:"enterprise_buyer,omitempty" xml:"enterprise_buyer,omitempty"`
}

var poolUserTagQueryResult = sync.Pool{
	New: func() any {
		return new(UserTagQueryResult)
	},
}

// GetUserTagQueryResult() 从对象池中获取UserTagQueryResult
func GetUserTagQueryResult() *UserTagQueryResult {
	return poolUserTagQueryResult.Get().(*UserTagQueryResult)
}

// ReleaseUserTagQueryResult 释放UserTagQueryResult
func ReleaseUserTagQueryResult(v *UserTagQueryResult) {
	v.EnterpriseBuyer = false
	poolUserTagQueryResult.Put(v)
}
