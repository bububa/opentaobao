package tbk

import (
	"sync"
)

// UpdateStatusResult 结构体
type UpdateStatusResult struct {
	// 暂停成功
	UpdateSuccess bool `json:"update_success,omitempty" xml:"update_success,omitempty"`
}

var poolUpdateStatusResult = sync.Pool{
	New: func() any {
		return new(UpdateStatusResult)
	},
}

// GetUpdateStatusResult() 从对象池中获取UpdateStatusResult
func GetUpdateStatusResult() *UpdateStatusResult {
	return poolUpdateStatusResult.Get().(*UpdateStatusResult)
}

// ReleaseUpdateStatusResult 释放UpdateStatusResult
func ReleaseUpdateStatusResult(v *UpdateStatusResult) {
	v.UpdateSuccess = false
	poolUpdateStatusResult.Put(v)
}
