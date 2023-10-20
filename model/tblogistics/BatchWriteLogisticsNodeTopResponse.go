package tblogistics

import (
	"sync"
)

// BatchWriteLogisticsNodeTopResponse 结构体
type BatchWriteLogisticsNodeTopResponse struct {
	// true成功，false失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchWriteLogisticsNodeTopResponse = sync.Pool{
	New: func() any {
		return new(BatchWriteLogisticsNodeTopResponse)
	},
}

// GetBatchWriteLogisticsNodeTopResponse() 从对象池中获取BatchWriteLogisticsNodeTopResponse
func GetBatchWriteLogisticsNodeTopResponse() *BatchWriteLogisticsNodeTopResponse {
	return poolBatchWriteLogisticsNodeTopResponse.Get().(*BatchWriteLogisticsNodeTopResponse)
}

// ReleaseBatchWriteLogisticsNodeTopResponse 释放BatchWriteLogisticsNodeTopResponse
func ReleaseBatchWriteLogisticsNodeTopResponse(v *BatchWriteLogisticsNodeTopResponse) {
	v.Success = false
	poolBatchWriteLogisticsNodeTopResponse.Put(v)
}
