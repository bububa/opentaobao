package tvupadmin

import (
	"sync"
)

// YunosTvpubadminManageDialogEditTopResult 结构体
type YunosTvpubadminManageDialogEditTopResult struct {
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

var poolYunosTvpubadminManageDialogEditTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminManageDialogEditTopResult)
	},
}

// GetYunosTvpubadminManageDialogEditTopResult() 从对象池中获取YunosTvpubadminManageDialogEditTopResult
func GetYunosTvpubadminManageDialogEditTopResult() *YunosTvpubadminManageDialogEditTopResult {
	return poolYunosTvpubadminManageDialogEditTopResult.Get().(*YunosTvpubadminManageDialogEditTopResult)
}

// ReleaseYunosTvpubadminManageDialogEditTopResult 释放YunosTvpubadminManageDialogEditTopResult
func ReleaseYunosTvpubadminManageDialogEditTopResult(v *YunosTvpubadminManageDialogEditTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.RetCode = 0
	v.Object = false
	v.Success = false
	poolYunosTvpubadminManageDialogEditTopResult.Put(v)
}
