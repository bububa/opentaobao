package tmallsc

import (
	"sync"
)

// AlibabaMsfserviceWorkerQueryidResult 结构体
type AlibabaMsfserviceWorkerQueryidResult struct {
	// 失败原因
	SystemError string `json:"system_error,omitempty" xml:"system_error,omitempty"`
	// 失败原因
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 工人id
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMsfserviceWorkerQueryidResult = sync.Pool{
	New: func() any {
		return new(AlibabaMsfserviceWorkerQueryidResult)
	},
}

// GetAlibabaMsfserviceWorkerQueryidResult() 从对象池中获取AlibabaMsfserviceWorkerQueryidResult
func GetAlibabaMsfserviceWorkerQueryidResult() *AlibabaMsfserviceWorkerQueryidResult {
	return poolAlibabaMsfserviceWorkerQueryidResult.Get().(*AlibabaMsfserviceWorkerQueryidResult)
}

// ReleaseAlibabaMsfserviceWorkerQueryidResult 释放AlibabaMsfserviceWorkerQueryidResult
func ReleaseAlibabaMsfserviceWorkerQueryidResult(v *AlibabaMsfserviceWorkerQueryidResult) {
	v.SystemError = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Result = 0
	v.Success = false
	poolAlibabaMsfserviceWorkerQueryidResult.Put(v)
}
