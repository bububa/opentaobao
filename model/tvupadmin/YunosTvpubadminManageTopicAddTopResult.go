package tvupadmin

import (
	"sync"
)

// YunosTvpubadminManageTopicAddTopResult 结构体
type YunosTvpubadminManageTopicAddTopResult struct {
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

var poolYunosTvpubadminManageTopicAddTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicAddTopResult)
	},
}

// GetYunosTvpubadminManageTopicAddTopResult() 从对象池中获取YunosTvpubadminManageTopicAddTopResult
func GetYunosTvpubadminManageTopicAddTopResult() *YunosTvpubadminManageTopicAddTopResult {
	return poolYunosTvpubadminManageTopicAddTopResult.Get().(*YunosTvpubadminManageTopicAddTopResult)
}

// ReleaseYunosTvpubadminManageTopicAddTopResult 释放YunosTvpubadminManageTopicAddTopResult
func ReleaseYunosTvpubadminManageTopicAddTopResult(v *YunosTvpubadminManageTopicAddTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Object = ""
	v.RetCode = 0
	v.Success = false
	poolYunosTvpubadminManageTopicAddTopResult.Put(v)
}
