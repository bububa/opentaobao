package tmallsc

import (
	"sync"
)

// AlibabaMsfserviceAuditsUpdateResult 结构体
type AlibabaMsfserviceAuditsUpdateResult struct {
	// 失败原因
	SystemError string `json:"system_error,omitempty" xml:"system_error,omitempty"`
	// 失败原因
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMsfserviceAuditsUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaMsfserviceAuditsUpdateResult)
	},
}

// GetAlibabaMsfserviceAuditsUpdateResult() 从对象池中获取AlibabaMsfserviceAuditsUpdateResult
func GetAlibabaMsfserviceAuditsUpdateResult() *AlibabaMsfserviceAuditsUpdateResult {
	return poolAlibabaMsfserviceAuditsUpdateResult.Get().(*AlibabaMsfserviceAuditsUpdateResult)
}

// ReleaseAlibabaMsfserviceAuditsUpdateResult 释放AlibabaMsfserviceAuditsUpdateResult
func ReleaseAlibabaMsfserviceAuditsUpdateResult(v *AlibabaMsfserviceAuditsUpdateResult) {
	v.SystemError = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolAlibabaMsfserviceAuditsUpdateResult.Put(v)
}
