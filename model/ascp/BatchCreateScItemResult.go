package ascp

import (
	"sync"
)

// BatchCreateScItemResult 结构体
type BatchCreateScItemResult struct {
	// 结果明细
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

var poolBatchCreateScItemResult = sync.Pool{
	New: func() any {
		return new(BatchCreateScItemResult)
	},
}

// GetBatchCreateScItemResult() 从对象池中获取BatchCreateScItemResult
func GetBatchCreateScItemResult() *BatchCreateScItemResult {
	return poolBatchCreateScItemResult.Get().(*BatchCreateScItemResult)
}

// ReleaseBatchCreateScItemResult 释放BatchCreateScItemResult
func ReleaseBatchCreateScItemResult(v *BatchCreateScItemResult) {
	v.Detail = v.Detail[:0]
	v.Result = ""
	poolBatchCreateScItemResult.Put(v)
}
