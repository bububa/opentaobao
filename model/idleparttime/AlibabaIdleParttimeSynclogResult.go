package idleparttime

import (
	"sync"
)

// AlibabaIdleParttimeSynclogResult 结构体
type AlibabaIdleParttimeSynclogResult struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据结构
	Data *AlibabaIdleParttimeSynclogData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleParttimeSynclogResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleParttimeSynclogResult)
	},
}

// GetAlibabaIdleParttimeSynclogResult() 从对象池中获取AlibabaIdleParttimeSynclogResult
func GetAlibabaIdleParttimeSynclogResult() *AlibabaIdleParttimeSynclogResult {
	return poolAlibabaIdleParttimeSynclogResult.Get().(*AlibabaIdleParttimeSynclogResult)
}

// ReleaseAlibabaIdleParttimeSynclogResult 释放AlibabaIdleParttimeSynclogResult
func ReleaseAlibabaIdleParttimeSynclogResult(v *AlibabaIdleParttimeSynclogResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolAlibabaIdleParttimeSynclogResult.Put(v)
}
