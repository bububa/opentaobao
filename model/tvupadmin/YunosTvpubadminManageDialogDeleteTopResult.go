package tvupadmin

import (
	"sync"
)

// YunosTvpubadminManageDialogDeleteTopResult 结构体
type YunosTvpubadminManageDialogDeleteTopResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// retCode
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// object
	Object int64 `json:"object,omitempty" xml:"object,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolYunosTvpubadminManageDialogDeleteTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogDeleteTopResult)
	},
}

// GetYunosTvpubadminManageDialogDeleteTopResult() 从对象池中获取YunosTvpubadminManageDialogDeleteTopResult
func GetYunosTvpubadminManageDialogDeleteTopResult() *YunosTvpubadminManageDialogDeleteTopResult {
	return poolYunosTvpubadminManageDialogDeleteTopResult.Get().(*YunosTvpubadminManageDialogDeleteTopResult)
}

// ReleaseYunosTvpubadminManageDialogDeleteTopResult 释放YunosTvpubadminManageDialogDeleteTopResult
func ReleaseYunosTvpubadminManageDialogDeleteTopResult(v *YunosTvpubadminManageDialogDeleteTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.RetCode = 0
	v.Object = 0
	v.Success = false
	poolYunosTvpubadminManageDialogDeleteTopResult.Put(v)
}
