package util

import (
	"sync"
)

// QimenResult 结构体
type QimenResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolQimenResult = sync.Pool{
	New: func() any {
		return new(QimenResult)
	},
}

// GetQimenResult() 从对象池中获取QimenResult
func GetQimenResult() *QimenResult {
	return poolQimenResult.Get().(*QimenResult)
}

// ReleaseQimenResult 释放QimenResult
func ReleaseQimenResult(v *QimenResult) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.IsSuccess = false
	poolQimenResult.Put(v)
}
