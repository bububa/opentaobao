package icbuassurance

import (
	"sync"
)

// AssuranceAccountResult 结构体
type AssuranceAccountResult struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// value
	Value *AssuranceFlag `json:"value,omitempty" xml:"value,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAssuranceAccountResult = sync.Pool{
	New: func() any {
		return new(AssuranceAccountResult)
	},
}

// GetAssuranceAccountResult() 从对象池中获取AssuranceAccountResult
func GetAssuranceAccountResult() *AssuranceAccountResult {
	return poolAssuranceAccountResult.Get().(*AssuranceAccountResult)
}

// ReleaseAssuranceAccountResult 释放AssuranceAccountResult
func ReleaseAssuranceAccountResult(v *AssuranceAccountResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Value = nil
	v.Success = false
	poolAssuranceAccountResult.Put(v)
}
