package mos

import (
	"sync"
)

// AlibabaMosStoreGetcloudshelfversionResultDo 结构体
type AlibabaMosStoreGetcloudshelfversionResultDo struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosStoreGetcloudshelfversionResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosStoreGetcloudshelfversionResultDo)
	},
}

// GetAlibabaMosStoreGetcloudshelfversionResultDo() 从对象池中获取AlibabaMosStoreGetcloudshelfversionResultDo
func GetAlibabaMosStoreGetcloudshelfversionResultDo() *AlibabaMosStoreGetcloudshelfversionResultDo {
	return poolAlibabaMosStoreGetcloudshelfversionResultDo.Get().(*AlibabaMosStoreGetcloudshelfversionResultDo)
}

// ReleaseAlibabaMosStoreGetcloudshelfversionResultDo 释放AlibabaMosStoreGetcloudshelfversionResultDo
func ReleaseAlibabaMosStoreGetcloudshelfversionResultDo(v *AlibabaMosStoreGetcloudshelfversionResultDo) {
	v.ErrMsg = ""
	v.Data = 0
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosStoreGetcloudshelfversionResultDo.Put(v)
}
