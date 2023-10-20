package alilabs

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// ret_msg
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// ret_value
	RetValue *HotWordsContent `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// ret_code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

var poolBaseResult = sync.Pool{
	New: func() any {
		return new(BaseResult)
	},
}

// GetBaseResult() 从对象池中获取BaseResult
func GetBaseResult() *BaseResult {
	return poolBaseResult.Get().(*BaseResult)
}

// ReleaseBaseResult 释放BaseResult
func ReleaseBaseResult(v *BaseResult) {
	v.RetMsg = ""
	v.RetValue = nil
	v.RetCode = 0
	poolBaseResult.Put(v)
}
