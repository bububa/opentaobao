package drugtrace

import (
	"sync"
)

// Header 结构体
type Header struct {
	// 消息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 标记
	SuccessFlag string `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}

var poolHeader = sync.Pool{
	New: func() any {
		return new(Header)
	},
}

// GetHeader() 从对象池中获取Header
func GetHeader() *Header {
	return poolHeader.Get().(*Header)
}

// ReleaseHeader 释放Header
func ReleaseHeader(v *Header) {
	v.ErrMsg = ""
	v.SuccessFlag = ""
	v.ErrCode = ""
	poolHeader.Put(v)
}
