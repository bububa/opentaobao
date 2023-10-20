package blackvip

import (
	"sync"
)

// ResultSupport 结构体
type ResultSupport struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 用户数据
	Models *Models `json:"models,omitempty" xml:"models,omitempty"`
	// 结果是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultSupport = sync.Pool{
	New: func() any {
		return new(ResultSupport)
	},
}

// GetResultSupport() 从对象池中获取ResultSupport
func GetResultSupport() *ResultSupport {
	return poolResultSupport.Get().(*ResultSupport)
}

// ReleaseResultSupport 释放ResultSupport
func ReleaseResultSupport(v *ResultSupport) {
	v.ResultCode = ""
	v.Models = nil
	v.Success = false
	poolResultSupport.Put(v)
}
