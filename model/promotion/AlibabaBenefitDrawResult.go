package promotion

import (
	"sync"
)

// AlibabaBenefitDrawResult 结构体
type AlibabaBenefitDrawResult struct {
	// message
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

var poolAlibabaBenefitDrawResult = sync.Pool{
	New: func() any {
		return new(AlibabaBenefitDrawResult)
	},
}

// GetAlibabaBenefitDrawResult() 从对象池中获取AlibabaBenefitDrawResult
func GetAlibabaBenefitDrawResult() *AlibabaBenefitDrawResult {
	return poolAlibabaBenefitDrawResult.Get().(*AlibabaBenefitDrawResult)
}

// ReleaseAlibabaBenefitDrawResult 释放AlibabaBenefitDrawResult
func ReleaseAlibabaBenefitDrawResult(v *AlibabaBenefitDrawResult) {
	v.ResultMsg = ""
	v.ResultCode = ""
	v.ResultSuccess = false
	poolAlibabaBenefitDrawResult.Put(v)
}
