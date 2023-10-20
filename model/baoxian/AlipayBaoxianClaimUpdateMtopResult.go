package baoxian

import (
	"sync"
)

// AlipayBaoxianClaimUpdateMtopResult 结构体
type AlipayBaoxianClaimUpdateMtopResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlipayBaoxianClaimUpdateMtopResult = sync.Pool{
	New: func() any {
		return new(AlipayBaoxianClaimUpdateMtopResult)
	},
}

// GetAlipayBaoxianClaimUpdateMtopResult() 从对象池中获取AlipayBaoxianClaimUpdateMtopResult
func GetAlipayBaoxianClaimUpdateMtopResult() *AlipayBaoxianClaimUpdateMtopResult {
	return poolAlipayBaoxianClaimUpdateMtopResult.Get().(*AlipayBaoxianClaimUpdateMtopResult)
}

// ReleaseAlipayBaoxianClaimUpdateMtopResult 释放AlipayBaoxianClaimUpdateMtopResult
func ReleaseAlipayBaoxianClaimUpdateMtopResult(v *AlipayBaoxianClaimUpdateMtopResult) {
	v.Model = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.IsSuccess = false
	poolAlipayBaoxianClaimUpdateMtopResult.Put(v)
}
