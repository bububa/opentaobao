package btrip

import (
	"sync"
)

// HisvResult 结构体
type HisvResult struct {
	// 审批单列表
	ModuleList []OpenIsvApplyRs `json:"module_list,omitempty" xml:"module_list>open_isv_apply_rs,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	Module *OpenApiApplyRs `json:"module,omitempty" xml:"module,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 分页相关信息
	PageInfo *PageInfoRs `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// 成功标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 成功标识
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolHisvResult = sync.Pool{
	New: func() any {
		return new(HisvResult)
	},
}

// GetHisvResult() 从对象池中获取HisvResult
func GetHisvResult() *HisvResult {
	return poolHisvResult.Get().(*HisvResult)
}

// ReleaseHisvResult 释放HisvResult
func ReleaseHisvResult(v *HisvResult) {
	v.ModuleList = v.ModuleList[:0]
	v.ResultMsg = ""
	v.Module = nil
	v.ResultCode = 0
	v.PageInfo = nil
	v.Success = false
	v.IsSuccess = false
	poolHisvResult.Put(v)
}
