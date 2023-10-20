package ascp

import (
	"sync"
)

// UnbundleItemmappingResult 结构体
type UnbundleItemmappingResult struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolUnbundleItemmappingResult = sync.Pool{
	New: func() any {
		return new(UnbundleItemmappingResult)
	},
}

// GetUnbundleItemmappingResult() 从对象池中获取UnbundleItemmappingResult
func GetUnbundleItemmappingResult() *UnbundleItemmappingResult {
	return poolUnbundleItemmappingResult.Get().(*UnbundleItemmappingResult)
}

// ReleaseUnbundleItemmappingResult 释放UnbundleItemmappingResult
func ReleaseUnbundleItemmappingResult(v *UnbundleItemmappingResult) {
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolUnbundleItemmappingResult.Put(v)
}
