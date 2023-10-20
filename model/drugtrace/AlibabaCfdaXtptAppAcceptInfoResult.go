package drugtrace

import (
	"sync"
)

// AlibabaCfdaXtptAppAcceptInfoResult 结构体
type AlibabaCfdaXtptAppAcceptInfoResult struct {
	// 消息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaCfdaXtptAppAcceptInfoResult = sync.Pool{
	New: func() any {
		return new(AlibabaCfdaXtptAppAcceptInfoResult)
	},
}

// GetAlibabaCfdaXtptAppAcceptInfoResult() 从对象池中获取AlibabaCfdaXtptAppAcceptInfoResult
func GetAlibabaCfdaXtptAppAcceptInfoResult() *AlibabaCfdaXtptAppAcceptInfoResult {
	return poolAlibabaCfdaXtptAppAcceptInfoResult.Get().(*AlibabaCfdaXtptAppAcceptInfoResult)
}

// ReleaseAlibabaCfdaXtptAppAcceptInfoResult 释放AlibabaCfdaXtptAppAcceptInfoResult
func ReleaseAlibabaCfdaXtptAppAcceptInfoResult(v *AlibabaCfdaXtptAppAcceptInfoResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaCfdaXtptAppAcceptInfoResult.Put(v)
}
