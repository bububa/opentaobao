package ascp

import (
	"sync"
)

// BatchUpdateScItemResult 结构体
type BatchUpdateScItemResult struct {
	// 结果明细
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

var poolBatchUpdateScItemResult = sync.Pool{
	New: func() any {
		return new(BatchUpdateScItemResult)
	},
}

// GetBatchUpdateScItemResult() 从对象池中获取BatchUpdateScItemResult
func GetBatchUpdateScItemResult() *BatchUpdateScItemResult {
	return poolBatchUpdateScItemResult.Get().(*BatchUpdateScItemResult)
}

// ReleaseBatchUpdateScItemResult 释放BatchUpdateScItemResult
func ReleaseBatchUpdateScItemResult(v *BatchUpdateScItemResult) {
	v.Detail = v.Detail[:0]
	v.Result = ""
	poolBatchUpdateScItemResult.Put(v)
}
