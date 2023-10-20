package dt

import (
	"sync"
)

// NrsResult 结构体
type NrsResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回数据
	Data *RecongnizeItemInfo `json:"data,omitempty" xml:"data,omitempty"`
	// 接口调用标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolNrsResult = sync.Pool{
	New: func() any {
		return new(NrsResult)
	},
}

// GetNrsResult() 从对象池中获取NrsResult
func GetNrsResult() *NrsResult {
	return poolNrsResult.Get().(*NrsResult)
}

// ReleaseNrsResult 释放NrsResult
func ReleaseNrsResult(v *NrsResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	v.Success = false
	poolNrsResult.Put(v)
}
