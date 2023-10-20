package inventory

import (
	"sync"
)

// BatchResult 结构体
type BatchResult struct {
	// 成功列表
	SuccessResultMap string `json:"success_result_map,omitempty" xml:"success_result_map,omitempty"`
	// 失败列表
	ErrorResultMap string `json:"error_result_map,omitempty" xml:"error_result_map,omitempty"`
	// 有一个失败，则整体是失败
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolBatchResult = sync.Pool{
	New: func() any {
		return new(BatchResult)
	},
}

// GetBatchResult() 从对象池中获取BatchResult
func GetBatchResult() *BatchResult {
	return poolBatchResult.Get().(*BatchResult)
}

// ReleaseBatchResult 释放BatchResult
func ReleaseBatchResult(v *BatchResult) {
	v.SuccessResultMap = ""
	v.ErrorResultMap = ""
	v.ResultCode = nil
	poolBatchResult.Put(v)
}
