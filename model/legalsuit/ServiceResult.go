package legalsuit

import (
	"sync"
)

// ServiceResult 结构体
type ServiceResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 案件内容
	Content *CaseModel `json:"content,omitempty" xml:"content,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolServiceResult = sync.Pool{
	New: func() any {
		return new(ServiceResult)
	},
}

// GetServiceResult() 从对象池中获取ServiceResult
func GetServiceResult() *ServiceResult {
	return poolServiceResult.Get().(*ServiceResult)
}

// ReleaseServiceResult 释放ServiceResult
func ReleaseServiceResult(v *ServiceResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolServiceResult.Put(v)
}
