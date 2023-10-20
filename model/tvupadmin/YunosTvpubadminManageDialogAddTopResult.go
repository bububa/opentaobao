package tvupadmin

import (
	"sync"
)

// YunosTvpubadminManageDialogAddTopResult 结构体
type YunosTvpubadminManageDialogAddTopResult struct {
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

var poolYunosTvpubadminManageDialogAddTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogAddTopResult)
	},
}

// GetYunosTvpubadminManageDialogAddTopResult() 从对象池中获取YunosTvpubadminManageDialogAddTopResult
func GetYunosTvpubadminManageDialogAddTopResult() *YunosTvpubadminManageDialogAddTopResult {
	return poolYunosTvpubadminManageDialogAddTopResult.Get().(*YunosTvpubadminManageDialogAddTopResult)
}

// ReleaseYunosTvpubadminManageDialogAddTopResult 释放YunosTvpubadminManageDialogAddTopResult
func ReleaseYunosTvpubadminManageDialogAddTopResult(v *YunosTvpubadminManageDialogAddTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.RetCode = 0
	v.Object = false
	v.Success = false
	poolYunosTvpubadminManageDialogAddTopResult.Put(v)
}
