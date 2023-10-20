package security

import (
	"sync"
)

// RpEventResult 结构体
type RpEventResult struct {
	// verifyLimitedFlag
	VerifyLimitedFlag bool `json:"verify_limited_flag,omitempty" xml:"verify_limited_flag,omitempty"`
}

var poolRpEventResult = sync.Pool{
	New: func() any {
		return new(RpEventResult)
	},
}

// GetRpEventResult() 从对象池中获取RpEventResult
func GetRpEventResult() *RpEventResult {
	return poolRpEventResult.Get().(*RpEventResult)
}

// ReleaseRpEventResult 释放RpEventResult
func ReleaseRpEventResult(v *RpEventResult) {
	v.VerifyLimitedFlag = false
	poolRpEventResult.Put(v)
}
