package mos

import (
	"sync"
)

// AlibabaMosStoreGetstorelistResultDo 结构体
type AlibabaMosStoreGetstorelistResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *MjStoresTopVo `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosStoreGetstorelistResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosStoreGetstorelistResultDo)
	},
}

// GetAlibabaMosStoreGetstorelistResultDo() 从对象池中获取AlibabaMosStoreGetstorelistResultDo
func GetAlibabaMosStoreGetstorelistResultDo() *AlibabaMosStoreGetstorelistResultDo {
	return poolAlibabaMosStoreGetstorelistResultDo.Get().(*AlibabaMosStoreGetstorelistResultDo)
}

// ReleaseAlibabaMosStoreGetstorelistResultDo 释放AlibabaMosStoreGetstorelistResultDo
func ReleaseAlibabaMosStoreGetstorelistResultDo(v *AlibabaMosStoreGetstorelistResultDo) {
	v.ErrMsg = ""
	v.Data = nil
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosStoreGetstorelistResultDo.Put(v)
}
