package xhotelitem

import (
	"sync"
)

// HsfResult 结构体
type HsfResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Module *SellerStatSummaryResult `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolHsfResult = sync.Pool{
	New: func() any {
		return new(HsfResult)
	},
}

// GetHsfResult() 从对象池中获取HsfResult
func GetHsfResult() *HsfResult {
	return poolHsfResult.Get().(*HsfResult)
}

// ReleaseHsfResult 释放HsfResult
func ReleaseHsfResult(v *HsfResult) {
	v.Code = ""
	v.Message = ""
	v.Module = nil
	v.Success = false
	poolHsfResult.Put(v)
}
