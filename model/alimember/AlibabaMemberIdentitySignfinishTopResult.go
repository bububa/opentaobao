package alimember

import (
	"sync"
)

// AlibabaMemberIdentitySignfinishTopResult 结构体
type AlibabaMemberIdentitySignfinishTopResult struct {
	// code，返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message，返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaMemberIdentitySignfinishTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaMemberIdentitySignfinishTopResult)
	},
}

// GetAlibabaMemberIdentitySignfinishTopResult() 从对象池中获取AlibabaMemberIdentitySignfinishTopResult
func GetAlibabaMemberIdentitySignfinishTopResult() *AlibabaMemberIdentitySignfinishTopResult {
	return poolAlibabaMemberIdentitySignfinishTopResult.Get().(*AlibabaMemberIdentitySignfinishTopResult)
}

// ReleaseAlibabaMemberIdentitySignfinishTopResult 释放AlibabaMemberIdentitySignfinishTopResult
func ReleaseAlibabaMemberIdentitySignfinishTopResult(v *AlibabaMemberIdentitySignfinishTopResult) {
	v.Code = ""
	v.Message = ""
	poolAlibabaMemberIdentitySignfinishTopResult.Put(v)
}
