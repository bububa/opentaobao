package tmc

import (
	"sync"
)

// TmcProduceResult 结构体
type TmcProduceResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolTmcProduceResult = sync.Pool{
	New: func() any {
		return new(TmcProduceResult)
	},
}

// GetTmcProduceResult() 从对象池中获取TmcProduceResult
func GetTmcProduceResult() *TmcProduceResult {
	return poolTmcProduceResult.Get().(*TmcProduceResult)
}

// ReleaseTmcProduceResult 释放TmcProduceResult
func ReleaseTmcProduceResult(v *TmcProduceResult) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.IsSuccess = false
	poolTmcProduceResult.Put(v)
}
