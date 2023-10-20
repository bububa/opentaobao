package ascp

import (
	"sync"
)

// BatchCreateCombineScItemResult 结构体
type BatchCreateCombineScItemResult struct {
	// 结果明细
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

var poolBatchCreateCombineScItemResult = sync.Pool{
	New: func() any {
		return new(BatchCreateCombineScItemResult)
	},
}

// GetBatchCreateCombineScItemResult() 从对象池中获取BatchCreateCombineScItemResult
func GetBatchCreateCombineScItemResult() *BatchCreateCombineScItemResult {
	return poolBatchCreateCombineScItemResult.Get().(*BatchCreateCombineScItemResult)
}

// ReleaseBatchCreateCombineScItemResult 释放BatchCreateCombineScItemResult
func ReleaseBatchCreateCombineScItemResult(v *BatchCreateCombineScItemResult) {
	v.Detail = v.Detail[:0]
	v.Result = ""
	poolBatchCreateCombineScItemResult.Put(v)
}
