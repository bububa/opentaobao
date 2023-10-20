package ascp

import (
	"sync"
)

// BatchUploadChannelInventoryResult 结构体
type BatchUploadChannelInventoryResult struct {
	// 结果明细
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

var poolBatchUploadChannelInventoryResult = sync.Pool{
	New: func() any {
		return new(BatchUploadChannelInventoryResult)
	},
}

// GetBatchUploadChannelInventoryResult() 从对象池中获取BatchUploadChannelInventoryResult
func GetBatchUploadChannelInventoryResult() *BatchUploadChannelInventoryResult {
	return poolBatchUploadChannelInventoryResult.Get().(*BatchUploadChannelInventoryResult)
}

// ReleaseBatchUploadChannelInventoryResult 释放BatchUploadChannelInventoryResult
func ReleaseBatchUploadChannelInventoryResult(v *BatchUploadChannelInventoryResult) {
	v.Detail = v.Detail[:0]
	v.Result = ""
	poolBatchUploadChannelInventoryResult.Put(v)
}
