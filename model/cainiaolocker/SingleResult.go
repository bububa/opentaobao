package cainiaolocker

import (
	"sync"
)

// SingleResult 结构体
type SingleResult struct {
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 参照返回码定义列表
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data *CainiaoEndpointLockerTopOrderNoticesendQueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolSingleResult = sync.Pool{
	New: func() any {
		return new(SingleResult)
	},
}

// GetSingleResult() 从对象池中获取SingleResult
func GetSingleResult() *SingleResult {
	return poolSingleResult.Get().(*SingleResult)
}

// ReleaseSingleResult 释放SingleResult
func ReleaseSingleResult(v *SingleResult) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolSingleResult.Put(v)
}
