package moscm

import (
	"sync"
)

// BulkInputCspuResult 结构体
type BulkInputCspuResult struct {
	// 录入结果对象
	CspuResultList []InputCspuResult `json:"cspu_result_list,omitempty" xml:"cspu_result_list>input_cspu_result,omitempty"`
}

var poolBulkInputCspuResult = sync.Pool{
	New: func() any {
		return new(BulkInputCspuResult)
	},
}

// GetBulkInputCspuResult() 从对象池中获取BulkInputCspuResult
func GetBulkInputCspuResult() *BulkInputCspuResult {
	return poolBulkInputCspuResult.Get().(*BulkInputCspuResult)
}

// ReleaseBulkInputCspuResult 释放BulkInputCspuResult
func ReleaseBulkInputCspuResult(v *BulkInputCspuResult) {
	v.CspuResultList = v.CspuResultList[:0]
	poolBulkInputCspuResult.Put(v)
}
