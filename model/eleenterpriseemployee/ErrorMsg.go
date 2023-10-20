package eleenterpriseemployee

import (
	"sync"
)

// ErrorMsg 结构体
type ErrorMsg struct {
	// 请求报文
	ReqBody string `json:"req_body,omitempty" xml:"req_body,omitempty"`
	// 失败原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
}

var poolErrorMsg = sync.Pool{
	New: func() any {
		return new(ErrorMsg)
	},
}

// GetErrorMsg() 从对象池中获取ErrorMsg
func GetErrorMsg() *ErrorMsg {
	return poolErrorMsg.Get().(*ErrorMsg)
}

// ReleaseErrorMsg 释放ErrorMsg
func ReleaseErrorMsg(v *ErrorMsg) {
	v.ReqBody = ""
	v.Reason = ""
	poolErrorMsg.Put(v)
}
