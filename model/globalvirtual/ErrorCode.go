package globalvirtual

import (
	"sync"
)

// ErrorCode 结构体
type ErrorCode struct {
	// error code key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// error code display message
	DisplayMessage string `json:"display_message,omitempty" xml:"display_message,omitempty"`
}

var poolErrorCode = sync.Pool{
	New: func() any {
		return new(ErrorCode)
	},
}

// GetErrorCode() 从对象池中获取ErrorCode
func GetErrorCode() *ErrorCode {
	return poolErrorCode.Get().(*ErrorCode)
}

// ReleaseErrorCode 释放ErrorCode
func ReleaseErrorCode(v *ErrorCode) {
	v.Key = ""
	v.DisplayMessage = ""
	poolErrorCode.Put(v)
}
