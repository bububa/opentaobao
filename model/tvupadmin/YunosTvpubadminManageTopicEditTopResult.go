package tvupadmin

import (
	"sync"
)

// YunosTvpubadminManageTopicEditTopResult 结构体
type YunosTvpubadminManageTopicEditTopResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// retCode
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolYunosTvpubadminManageTopicEditTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicEditTopResult)
	},
}

// GetYunosTvpubadminManageTopicEditTopResult() 从对象池中获取YunosTvpubadminManageTopicEditTopResult
func GetYunosTvpubadminManageTopicEditTopResult() *YunosTvpubadminManageTopicEditTopResult {
	return poolYunosTvpubadminManageTopicEditTopResult.Get().(*YunosTvpubadminManageTopicEditTopResult)
}

// ReleaseYunosTvpubadminManageTopicEditTopResult 释放YunosTvpubadminManageTopicEditTopResult
func ReleaseYunosTvpubadminManageTopicEditTopResult(v *YunosTvpubadminManageTopicEditTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.RetCode = 0
	v.Object = false
	v.Success = false
	poolYunosTvpubadminManageTopicEditTopResult.Put(v)
}
