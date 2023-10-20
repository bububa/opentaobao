package xhotelitem

import (
	"sync"
)

// BnbResult 结构体
type BnbResult struct {
	// 响应码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 响应信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 状态，成功true，失败false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBnbResult = sync.Pool{
	New: func() any {
		return new(BnbResult)
	},
}

// GetBnbResult() 从对象池中获取BnbResult
func GetBnbResult() *BnbResult {
	return poolBnbResult.Get().(*BnbResult)
}

// ReleaseBnbResult 释放BnbResult
func ReleaseBnbResult(v *BnbResult) {
	v.ResultCode = ""
	v.ResultMsg = ""
	v.Success = false
	poolBnbResult.Put(v)
}
