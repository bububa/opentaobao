package tvupadmin

import (
	"sync"
)

// YunosTvpubadminManageTopicContentaddTopResult 结构体
type YunosTvpubadminManageTopicContentaddTopResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// retCode
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolYunosTvpubadminManageTopicContentaddTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicContentaddTopResult)
	},
}

// GetYunosTvpubadminManageTopicContentaddTopResult() 从对象池中获取YunosTvpubadminManageTopicContentaddTopResult
func GetYunosTvpubadminManageTopicContentaddTopResult() *YunosTvpubadminManageTopicContentaddTopResult {
	return poolYunosTvpubadminManageTopicContentaddTopResult.Get().(*YunosTvpubadminManageTopicContentaddTopResult)
}

// ReleaseYunosTvpubadminManageTopicContentaddTopResult 释放YunosTvpubadminManageTopicContentaddTopResult
func ReleaseYunosTvpubadminManageTopicContentaddTopResult(v *YunosTvpubadminManageTopicContentaddTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Object = ""
	v.RetCode = 0
	v.Success = false
	poolYunosTvpubadminManageTopicContentaddTopResult.Put(v)
}
