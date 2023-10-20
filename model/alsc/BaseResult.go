package alsc

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1成功，2失败，3处理中
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否失败
	Fail bool `json:"fail,omitempty" xml:"fail,omitempty"`
	// 是否处理中
	Process bool `json:"process,omitempty" xml:"process,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBaseResult = sync.Pool{
	New: func() any {
		return new(BaseResult)
	},
}

// GetBaseResult() 从对象池中获取BaseResult
func GetBaseResult() *BaseResult {
	return poolBaseResult.Get().(*BaseResult)
}

// ReleaseBaseResult 释放BaseResult
func ReleaseBaseResult(v *BaseResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Status = 0
	v.Fail = false
	v.Process = false
	v.Success = false
	poolBaseResult.Put(v)
}
