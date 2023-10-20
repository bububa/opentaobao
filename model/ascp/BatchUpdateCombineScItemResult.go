package ascp

import (
	"sync"
)

// BatchUpdateCombineScItemResult 结构体
type BatchUpdateCombineScItemResult struct {
	// 结果明细
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

var poolBatchUpdateCombineScItemResult = sync.Pool{
	New: func() any {
		return new(BatchUpdateCombineScItemResult)
	},
}

// GetBatchUpdateCombineScItemResult() 从对象池中获取BatchUpdateCombineScItemResult
func GetBatchUpdateCombineScItemResult() *BatchUpdateCombineScItemResult {
	return poolBatchUpdateCombineScItemResult.Get().(*BatchUpdateCombineScItemResult)
}

// ReleaseBatchUpdateCombineScItemResult 释放BatchUpdateCombineScItemResult
func ReleaseBatchUpdateCombineScItemResult(v *BatchUpdateCombineScItemResult) {
	v.Detail = v.Detail[:0]
	v.Result = ""
	poolBatchUpdateCombineScItemResult.Put(v)
}
