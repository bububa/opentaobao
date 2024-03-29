package btrip

import (
	"sync"
)

// BtmsResult 结构体
type BtmsResult struct {
	// 结果描述。
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 分页结果对象。
	Module *PagingResult `json:"module,omitempty" xml:"module,omitempty"`
	// 结果码；0：成功，非0：失败。
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 请求是否成功。
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBtmsResult = sync.Pool{
	New: func() any {
		return new(BtmsResult)
	},
}

// GetBtmsResult() 从对象池中获取BtmsResult
func GetBtmsResult() *BtmsResult {
	return poolBtmsResult.Get().(*BtmsResult)
}

// ReleaseBtmsResult 释放BtmsResult
func ReleaseBtmsResult(v *BtmsResult) {
	v.ResultMsg = ""
	v.Module = nil
	v.ResultCode = 0
	v.Success = false
	poolBtmsResult.Put(v)
}
