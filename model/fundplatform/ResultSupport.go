package fundplatform

import (
	"sync"
)

// ResultSupport 结构体
type ResultSupport struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 出参对象
	Module *AccountChargeResponse `json:"module,omitempty" xml:"module,omitempty"`
	// success
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
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Module = nil
	v.Success = false
	poolResultSupport.Put(v)
}
