package ascp

import (
	"sync"
)

// BatchCreateItemMappingResult 结构体
type BatchCreateItemMappingResult struct {
	// 结果明细
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

var poolBatchCreateItemMappingResult = sync.Pool{
	New: func() any {
		return new(BatchCreateItemMappingResult)
	},
}

// GetBatchCreateItemMappingResult() 从对象池中获取BatchCreateItemMappingResult
func GetBatchCreateItemMappingResult() *BatchCreateItemMappingResult {
	return poolBatchCreateItemMappingResult.Get().(*BatchCreateItemMappingResult)
}

// ReleaseBatchCreateItemMappingResult 释放BatchCreateItemMappingResult
func ReleaseBatchCreateItemMappingResult(v *BatchCreateItemMappingResult) {
	v.Detail = v.Detail[:0]
	v.Result = ""
	poolBatchCreateItemMappingResult.Put(v)
}
