package yunosminiapp

import (
	"sync"
)

// TopActivityResult 结构体
type TopActivityResult struct {
	// 详细信息
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 成功与否
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 活动状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolTopActivityResult = sync.Pool{
	New: func() any {
		return new(TopActivityResult)
	},
}

// GetTopActivityResult() 从对象池中获取TopActivityResult
func GetTopActivityResult() *TopActivityResult {
	return poolTopActivityResult.Get().(*TopActivityResult)
}

// ReleaseTopActivityResult 释放TopActivityResult
func ReleaseTopActivityResult(v *TopActivityResult) {
	v.Detail = ""
	v.Success = ""
	v.Status = ""
	poolTopActivityResult.Put(v)
}
