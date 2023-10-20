package ascp

import (
	"sync"
)

// BatchQueryInventoryResult 结构体
type BatchQueryInventoryResult struct {
	// 结果明细
	Detail []ScItemInfo `json:"detail,omitempty" xml:"detail>sc_item_info,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchQueryInventoryResult = sync.Pool{
	New: func() any {
		return new(BatchQueryInventoryResult)
	},
}

// GetBatchQueryInventoryResult() 从对象池中获取BatchQueryInventoryResult
func GetBatchQueryInventoryResult() *BatchQueryInventoryResult {
	return poolBatchQueryInventoryResult.Get().(*BatchQueryInventoryResult)
}

// ReleaseBatchQueryInventoryResult 释放BatchQueryInventoryResult
func ReleaseBatchQueryInventoryResult(v *BatchQueryInventoryResult) {
	v.Detail = v.Detail[:0]
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolBatchQueryInventoryResult.Put(v)
}
