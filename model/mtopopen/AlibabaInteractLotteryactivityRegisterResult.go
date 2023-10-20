package mtopopen

import (
	"sync"
)

// AlibabaInteractLotteryactivityRegisterResult 结构体
type AlibabaInteractLotteryactivityRegisterResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 方便追踪请求的唯一标识
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回的数据
	Data *ActivityLotteryWriteResult `json:"data,omitempty" xml:"data,omitempty"`
	// xxx
	ModuleMap *Modulemap `json:"module_map,omitempty" xml:"module_map,omitempty"`
	// 注册抽奖活动失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaInteractLotteryactivityRegisterResult = sync.Pool{
	New: func() any {
		return new(AlibabaInteractLotteryactivityRegisterResult)
	},
}

// GetAlibabaInteractLotteryactivityRegisterResult() 从对象池中获取AlibabaInteractLotteryactivityRegisterResult
func GetAlibabaInteractLotteryactivityRegisterResult() *AlibabaInteractLotteryactivityRegisterResult {
	return poolAlibabaInteractLotteryactivityRegisterResult.Get().(*AlibabaInteractLotteryactivityRegisterResult)
}

// ReleaseAlibabaInteractLotteryactivityRegisterResult 释放AlibabaInteractLotteryactivityRegisterResult
func ReleaseAlibabaInteractLotteryactivityRegisterResult(v *AlibabaInteractLotteryactivityRegisterResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.TraceId = ""
	v.Data = nil
	v.ModuleMap = nil
	v.Success = false
	poolAlibabaInteractLotteryactivityRegisterResult.Put(v)
}
