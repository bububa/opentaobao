package alihealthcert

import (
	"sync"
)

// ServiceResult 结构体
type ServiceResult struct {
	// 111
	EagleEyeTraceId string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
	// 系统错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 系统错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 返回数据对象
	Data *ReserveStatusResultResponse `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.EagleEyeTraceId = ""
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Data = nil
	v.Success = false
	poolServiceResult.Put(v)
}
