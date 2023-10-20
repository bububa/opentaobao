package tmallhk

import (
	"sync"
)

// DataResult 结构体
type DataResult struct {
	// 参数code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 参数msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// obj
	Obj *CCICCheckCodeDo `json:"obj,omitempty" xml:"obj,omitempty"`
	// 是否正常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDataResult = sync.Pool{
	New: func() any {
		return new(DataResult)
	},
}

// GetDataResult() 从对象池中获取DataResult
func GetDataResult() *DataResult {
	return poolDataResult.Get().(*DataResult)
}

// ReleaseDataResult 释放DataResult
func ReleaseDataResult(v *DataResult) {
	v.Code = ""
	v.Msg = ""
	v.Obj = nil
	v.Success = false
	poolDataResult.Put(v)
}
