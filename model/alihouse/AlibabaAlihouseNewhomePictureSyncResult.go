package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomePictureSyncResult 结构体
type AlibabaAlihouseNewhomePictureSyncResult struct {
	// 图片外部ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomePictureSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomePictureSyncResult)
	},
}

// GetAlibabaAlihouseNewhomePictureSyncResult() 从对象池中获取AlibabaAlihouseNewhomePictureSyncResult
func GetAlibabaAlihouseNewhomePictureSyncResult() *AlibabaAlihouseNewhomePictureSyncResult {
	return poolAlibabaAlihouseNewhomePictureSyncResult.Get().(*AlibabaAlihouseNewhomePictureSyncResult)
}

// ReleaseAlibabaAlihouseNewhomePictureSyncResult 释放AlibabaAlihouseNewhomePictureSyncResult
func ReleaseAlibabaAlihouseNewhomePictureSyncResult(v *AlibabaAlihouseNewhomePictureSyncResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomePictureSyncResult.Put(v)
}
