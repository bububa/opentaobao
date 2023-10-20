package btrip

import (
	"sync"
)

// BtripApplyResult 结构体
type BtripApplyResult struct {
	// 审批单列表
	ApplyList []OpenApplyRs `json:"apply_list,omitempty" xml:"apply_list>open_apply_rs,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果对象
	Module *OpenApiNewApplyRs `json:"module,omitempty" xml:"module,omitempty"`
	// 结果码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBtripApplyResult = sync.Pool{
	New: func() any {
		return new(BtripApplyResult)
	},
}

// GetBtripApplyResult() 从对象池中获取BtripApplyResult
func GetBtripApplyResult() *BtripApplyResult {
	return poolBtripApplyResult.Get().(*BtripApplyResult)
}

// ReleaseBtripApplyResult 释放BtripApplyResult
func ReleaseBtripApplyResult(v *BtripApplyResult) {
	v.ApplyList = v.ApplyList[:0]
	v.ResultMsg = ""
	v.ErrMsg = ""
	v.ResultCode = 0
	v.Module = nil
	v.ErrCode = 0
	v.Total = 0
	v.Success = false
	poolBtripApplyResult.Put(v)
}
