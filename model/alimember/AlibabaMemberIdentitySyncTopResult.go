package alimember

import (
	"sync"
)

// AlibabaMemberIdentitySyncTopResult 结构体
type AlibabaMemberIdentitySyncTopResult struct {
	// code，返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message，返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaMemberIdentitySyncTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaMemberIdentitySyncTopResult)
	},
}

// GetAlibabaMemberIdentitySyncTopResult() 从对象池中获取AlibabaMemberIdentitySyncTopResult
func GetAlibabaMemberIdentitySyncTopResult() *AlibabaMemberIdentitySyncTopResult {
	return poolAlibabaMemberIdentitySyncTopResult.Get().(*AlibabaMemberIdentitySyncTopResult)
}

// ReleaseAlibabaMemberIdentitySyncTopResult 释放AlibabaMemberIdentitySyncTopResult
func ReleaseAlibabaMemberIdentitySyncTopResult(v *AlibabaMemberIdentitySyncTopResult) {
	v.Code = ""
	v.Message = ""
	poolAlibabaMemberIdentitySyncTopResult.Put(v)
}
