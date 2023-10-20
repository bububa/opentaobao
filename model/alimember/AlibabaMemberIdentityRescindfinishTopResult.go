package alimember

import (
	"sync"
)

// AlibabaMemberIdentityRescindfinishTopResult 结构体
type AlibabaMemberIdentityRescindfinishTopResult struct {
	// code，返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message，返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaMemberIdentityRescindfinishTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaMemberIdentityRescindfinishTopResult)
	},
}

// GetAlibabaMemberIdentityRescindfinishTopResult() 从对象池中获取AlibabaMemberIdentityRescindfinishTopResult
func GetAlibabaMemberIdentityRescindfinishTopResult() *AlibabaMemberIdentityRescindfinishTopResult {
	return poolAlibabaMemberIdentityRescindfinishTopResult.Get().(*AlibabaMemberIdentityRescindfinishTopResult)
}

// ReleaseAlibabaMemberIdentityRescindfinishTopResult 释放AlibabaMemberIdentityRescindfinishTopResult
func ReleaseAlibabaMemberIdentityRescindfinishTopResult(v *AlibabaMemberIdentityRescindfinishTopResult) {
	v.Code = ""
	v.Message = ""
	poolAlibabaMemberIdentityRescindfinishTopResult.Put(v)
}
