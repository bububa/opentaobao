package customizemarket

import (
	"sync"
)

// ServiceResult 结构体
type ServiceResult struct {
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// suc
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
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
	v.Data = ""
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Suc = false
	poolServiceResult.Put(v)
}
