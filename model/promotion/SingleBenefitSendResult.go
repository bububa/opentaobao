package promotion

import (
	"sync"
)

// SingleBenefitSendResult 结构体
type SingleBenefitSendResult struct {
	// 单次发放结果
	Results []MobileBenefitSendResultExt `json:"results,omitempty" xml:"results>mobile_benefit_send_result_ext,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 事务id
	UniqueId string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
	// 是否发放成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 用于宝箱应用，为true表示宝箱任务已完成，需要提示用户
	IsBaoxiangTaskDone bool `json:"is_baoxiang_task_done,omitempty" xml:"is_baoxiang_task_done,omitempty"`
}

var poolSingleBenefitSendResult = sync.Pool{
	New: func() any {
		return new(SingleBenefitSendResult)
	},
}

// GetSingleBenefitSendResult() 从对象池中获取SingleBenefitSendResult
func GetSingleBenefitSendResult() *SingleBenefitSendResult {
	return poolSingleBenefitSendResult.Get().(*SingleBenefitSendResult)
}

// ReleaseSingleBenefitSendResult 释放SingleBenefitSendResult
func ReleaseSingleBenefitSendResult(v *SingleBenefitSendResult) {
	v.Results = v.Results[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.UniqueId = ""
	v.IsSuccess = false
	v.IsBaoxiangTaskDone = false
	poolSingleBenefitSendResult.Put(v)
}
