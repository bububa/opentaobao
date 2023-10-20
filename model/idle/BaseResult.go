package idle

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否发货成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否执行成功
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
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = false
	v.Success = false
	poolBaseResult.Put(v)
}
