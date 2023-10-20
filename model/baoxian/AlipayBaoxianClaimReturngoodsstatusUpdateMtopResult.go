package baoxian

import (
	"sync"
)

// AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult 结构体
type AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlipayBaoxianClaimReturngoodsstatusUpdateMtopResult = sync.Pool{
	New: func() any {
		return new(AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult)
	},
}

// GetAlipayBaoxianClaimReturngoodsstatusUpdateMtopResult() 从对象池中获取AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult
func GetAlipayBaoxianClaimReturngoodsstatusUpdateMtopResult() *AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult {
	return poolAlipayBaoxianClaimReturngoodsstatusUpdateMtopResult.Get().(*AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult)
}

// ReleaseAlipayBaoxianClaimReturngoodsstatusUpdateMtopResult 释放AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult
func ReleaseAlipayBaoxianClaimReturngoodsstatusUpdateMtopResult(v *AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult) {
	v.Model = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.IsSuccess = false
	poolAlipayBaoxianClaimReturngoodsstatusUpdateMtopResult.Put(v)
}
