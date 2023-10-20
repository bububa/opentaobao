package servicecenter

import (
	"sync"
)

// ErrorMessage 结构体
type ErrorMessage struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolErrorMessage = sync.Pool{
	New: func() any {
		return new(ErrorMessage)
	},
}

// GetErrorMessage() 从对象池中获取ErrorMessage
func GetErrorMessage() *ErrorMessage {
	return poolErrorMessage.Get().(*ErrorMessage)
}

// ReleaseErrorMessage 释放ErrorMessage
func ReleaseErrorMessage(v *ErrorMessage) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	poolErrorMessage.Put(v)
}
