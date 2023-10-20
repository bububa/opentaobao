package alimember

import (
	"sync"
)

// AlibabaMemberSyncResult 结构体
type AlibabaMemberSyncResult struct {
	// code，返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message，返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaMemberSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaMemberSyncResult)
	},
}

// GetAlibabaMemberSyncResult() 从对象池中获取AlibabaMemberSyncResult
func GetAlibabaMemberSyncResult() *AlibabaMemberSyncResult {
	return poolAlibabaMemberSyncResult.Get().(*AlibabaMemberSyncResult)
}

// ReleaseAlibabaMemberSyncResult 释放AlibabaMemberSyncResult
func ReleaseAlibabaMemberSyncResult(v *AlibabaMemberSyncResult) {
	v.Code = ""
	v.Message = ""
	poolAlibabaMemberSyncResult.Put(v)
}
