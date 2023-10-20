package alisports

import (
	"sync"
)

// AlispResult 结构体
type AlispResult struct {
	// 错误信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 数据对象
	AlispData *AuthAccountInfoDto `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// 错误码
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

var poolAlispResult = sync.Pool{
	New: func() any {
		return new(AlispResult)
	},
}

// GetAlispResult() 从对象池中获取AlispResult
func GetAlispResult() *AlispResult {
	return poolAlispResult.Get().(*AlispResult)
}

// ReleaseAlispResult 释放AlispResult
func ReleaseAlispResult(v *AlispResult) {
	v.AlispMsg = ""
	v.AlispData = nil
	v.AlispCode = 0
	poolAlispResult.Put(v)
}
