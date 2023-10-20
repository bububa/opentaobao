package mos

import (
	"sync"
)

// AlibabaMosOcTradeSyncbanksaleResultDo 结构体
type AlibabaMosOcTradeSyncbanksaleResultDo struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosOcTradeSyncbanksaleResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosOcTradeSyncbanksaleResultDo)
	},
}

// GetAlibabaMosOcTradeSyncbanksaleResultDo() 从对象池中获取AlibabaMosOcTradeSyncbanksaleResultDo
func GetAlibabaMosOcTradeSyncbanksaleResultDo() *AlibabaMosOcTradeSyncbanksaleResultDo {
	return poolAlibabaMosOcTradeSyncbanksaleResultDo.Get().(*AlibabaMosOcTradeSyncbanksaleResultDo)
}

// ReleaseAlibabaMosOcTradeSyncbanksaleResultDo 释放AlibabaMosOcTradeSyncbanksaleResultDo
func ReleaseAlibabaMosOcTradeSyncbanksaleResultDo(v *AlibabaMosOcTradeSyncbanksaleResultDo) {
	v.ErrMsg = ""
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosOcTradeSyncbanksaleResultDo.Put(v)
}
