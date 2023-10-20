package alilabs

import (
	"sync"
)

// AlibabaAilabUserProfileGetResult 结构体
type AlibabaAilabUserProfileGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 数据model
	Result *BasicUserInfo `json:"result,omitempty" xml:"result,omitempty"`
	// 200 成功，其他 失败
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

var poolAlibabaAilabUserProfileGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabUserProfileGetResult)
	},
}

// GetAlibabaAilabUserProfileGetResult() 从对象池中获取AlibabaAilabUserProfileGetResult
func GetAlibabaAilabUserProfileGetResult() *AlibabaAilabUserProfileGetResult {
	return poolAlibabaAilabUserProfileGetResult.Get().(*AlibabaAilabUserProfileGetResult)
}

// ReleaseAlibabaAilabUserProfileGetResult 释放AlibabaAilabUserProfileGetResult
func ReleaseAlibabaAilabUserProfileGetResult(v *AlibabaAilabUserProfileGetResult) {
	v.Message = ""
	v.Result = nil
	v.StatusCode = 0
	poolAlibabaAilabUserProfileGetResult.Put(v)
}
