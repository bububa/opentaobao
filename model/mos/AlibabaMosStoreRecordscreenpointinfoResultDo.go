package mos

import (
	"sync"
)

// AlibabaMosStoreRecordscreenpointinfoResultDo 结构体
type AlibabaMosStoreRecordscreenpointinfoResultDo struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosStoreRecordscreenpointinfoResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosStoreRecordscreenpointinfoResultDo)
	},
}

// GetAlibabaMosStoreRecordscreenpointinfoResultDo() 从对象池中获取AlibabaMosStoreRecordscreenpointinfoResultDo
func GetAlibabaMosStoreRecordscreenpointinfoResultDo() *AlibabaMosStoreRecordscreenpointinfoResultDo {
	return poolAlibabaMosStoreRecordscreenpointinfoResultDo.Get().(*AlibabaMosStoreRecordscreenpointinfoResultDo)
}

// ReleaseAlibabaMosStoreRecordscreenpointinfoResultDo 释放AlibabaMosStoreRecordscreenpointinfoResultDo
func ReleaseAlibabaMosStoreRecordscreenpointinfoResultDo(v *AlibabaMosStoreRecordscreenpointinfoResultDo) {
	v.ErrMsg = ""
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosStoreRecordscreenpointinfoResultDo.Put(v)
}
