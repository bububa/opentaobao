package idle

import (
	"sync"
)

// AlibabaIdleRecycleOrderFulfillmentResult 结构体
type AlibabaIdleRecycleOrderFulfillmentResult struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRecycleOrderFulfillmentResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleOrderFulfillmentResult)
	},
}

// GetAlibabaIdleRecycleOrderFulfillmentResult() 从对象池中获取AlibabaIdleRecycleOrderFulfillmentResult
func GetAlibabaIdleRecycleOrderFulfillmentResult() *AlibabaIdleRecycleOrderFulfillmentResult {
	return poolAlibabaIdleRecycleOrderFulfillmentResult.Get().(*AlibabaIdleRecycleOrderFulfillmentResult)
}

// ReleaseAlibabaIdleRecycleOrderFulfillmentResult 释放AlibabaIdleRecycleOrderFulfillmentResult
func ReleaseAlibabaIdleRecycleOrderFulfillmentResult(v *AlibabaIdleRecycleOrderFulfillmentResult) {
	v.Success = false
	poolAlibabaIdleRecycleOrderFulfillmentResult.Put(v)
}
