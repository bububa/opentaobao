package wdk

import (
	"sync"
)

// UtmsResult 结构体
type UtmsResult struct {
	// list
	List []BomProcessDto `json:"list,omitempty" xml:"list>bom_process_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolUtmsResult = sync.Pool{
	New: func() any {
		return new(UtmsResult)
	},
}

// GetUtmsResult() 从对象池中获取UtmsResult
func GetUtmsResult() *UtmsResult {
	return poolUtmsResult.Get().(*UtmsResult)
}

// ReleaseUtmsResult 释放UtmsResult
func ReleaseUtmsResult(v *UtmsResult) {
	v.List = v.List[:0]
	v.Code = ""
	v.Msg = ""
	v.Success = false
	v.Model = false
	poolUtmsResult.Put(v)
}
