package tvupadmin

import (
	"sync"
)

// YunosTvpubadminManageTopicContentdeleteTopResult 结构体
type YunosTvpubadminManageTopicContentdeleteTopResult struct {
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

var poolYunosTvpubadminManageTopicContentdeleteTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageTopicContentdeleteTopResult)
	},
}

// GetYunosTvpubadminManageTopicContentdeleteTopResult() 从对象池中获取YunosTvpubadminManageTopicContentdeleteTopResult
func GetYunosTvpubadminManageTopicContentdeleteTopResult() *YunosTvpubadminManageTopicContentdeleteTopResult {
	return poolYunosTvpubadminManageTopicContentdeleteTopResult.Get().(*YunosTvpubadminManageTopicContentdeleteTopResult)
}

// ReleaseYunosTvpubadminManageTopicContentdeleteTopResult 释放YunosTvpubadminManageTopicContentdeleteTopResult
func ReleaseYunosTvpubadminManageTopicContentdeleteTopResult(v *YunosTvpubadminManageTopicContentdeleteTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Object = ""
	v.RetCode = 0
	v.Success = false
	poolYunosTvpubadminManageTopicContentdeleteTopResult.Put(v)
}
