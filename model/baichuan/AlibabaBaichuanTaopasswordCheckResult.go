package baichuan

import (
	"sync"
)

// AlibabaBaichuanTaopasswordCheckResult 结构体
type AlibabaBaichuanTaopasswordCheckResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaBaichuanTaopasswordCheckResult = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanTaopasswordCheckResult)
	},
}

// GetAlibabaBaichuanTaopasswordCheckResult() 从对象池中获取AlibabaBaichuanTaopasswordCheckResult
func GetAlibabaBaichuanTaopasswordCheckResult() *AlibabaBaichuanTaopasswordCheckResult {
	return poolAlibabaBaichuanTaopasswordCheckResult.Get().(*AlibabaBaichuanTaopasswordCheckResult)
}

// ReleaseAlibabaBaichuanTaopasswordCheckResult 释放AlibabaBaichuanTaopasswordCheckResult
func ReleaseAlibabaBaichuanTaopasswordCheckResult(v *AlibabaBaichuanTaopasswordCheckResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = ""
	v.Success = false
	poolAlibabaBaichuanTaopasswordCheckResult.Put(v)
}
