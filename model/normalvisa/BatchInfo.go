package normalvisa

import (
	"sync"
)

// BatchInfo 结构体
type BatchInfo struct {
	// 每批次的申请人id
	ApplyIds []string `json:"apply_ids,omitempty" xml:"apply_ids>string,omitempty"`
	// 批次id
	BatchId string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
}

var poolBatchInfo = sync.Pool{
	New: func() any {
		return new(BatchInfo)
	},
}

// GetBatchInfo() 从对象池中获取BatchInfo
func GetBatchInfo() *BatchInfo {
	return poolBatchInfo.Get().(*BatchInfo)
}

// ReleaseBatchInfo 释放BatchInfo
func ReleaseBatchInfo(v *BatchInfo) {
	v.ApplyIds = v.ApplyIds[:0]
	v.BatchId = ""
	poolBatchInfo.Put(v)
}
