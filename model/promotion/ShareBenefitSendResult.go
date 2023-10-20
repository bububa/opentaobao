package promotion

import (
	"sync"
)

// ShareBenefitSendResult 结构体
type ShareBenefitSendResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 发放结果
	ResultMap string `json:"result_map,omitempty" xml:"result_map,omitempty"`
	// 发放结果是否正常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolShareBenefitSendResult = sync.Pool{
	New: func() any {
		return new(ShareBenefitSendResult)
	},
}

// GetShareBenefitSendResult() 从对象池中获取ShareBenefitSendResult
func GetShareBenefitSendResult() *ShareBenefitSendResult {
	return poolShareBenefitSendResult.Get().(*ShareBenefitSendResult)
}

// ReleaseShareBenefitSendResult 释放ShareBenefitSendResult
func ReleaseShareBenefitSendResult(v *ShareBenefitSendResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.ResultMap = ""
	v.Success = false
	poolShareBenefitSendResult.Put(v)
}
