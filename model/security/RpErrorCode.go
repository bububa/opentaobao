package security

import (
	"sync"
)

// RpErrorCode 结构体
type RpErrorCode struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorName
	ErrorName string `json:"error_name,omitempty" xml:"error_name,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
}

var poolRpErrorCode = sync.Pool{
	New: func() any {
		return new(RpErrorCode)
	},
}

// GetRpErrorCode() 从对象池中获取RpErrorCode
func GetRpErrorCode() *RpErrorCode {
	return poolRpErrorCode.Get().(*RpErrorCode)
}

// ReleaseRpErrorCode 释放RpErrorCode
func ReleaseRpErrorCode(v *RpErrorCode) {
	v.ErrorMsg = ""
	v.ErrorName = ""
	v.ErrorCode = 0
	poolRpErrorCode.Put(v)
}
