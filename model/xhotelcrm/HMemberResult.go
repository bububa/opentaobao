package xhotelcrm

import (
	"sync"
)

// HMemberResult 结构体
type HMemberResult struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 信息
	ResultInfo string `json:"result_info,omitempty" xml:"result_info,omitempty"`
	// 数据
	Module *MerchantMemberBindInfoVo `json:"module,omitempty" xml:"module,omitempty"`
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolHMemberResult = sync.Pool{
	New: func() any {
		return new(HMemberResult)
	},
}

// GetHMemberResult() 从对象池中获取HMemberResult
func GetHMemberResult() *HMemberResult {
	return poolHMemberResult.Get().(*HMemberResult)
}

// ReleaseHMemberResult 释放HMemberResult
func ReleaseHMemberResult(v *HMemberResult) {
	v.ResultCode = ""
	v.ResultInfo = ""
	v.Module = nil
	v.Success = false
	poolHMemberResult.Put(v)
}
