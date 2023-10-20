package mos

import (
	"sync"
)

// AlibabaMosStoreGetdefautitemsResultDo 结构体
type AlibabaMosStoreGetdefautitemsResultDo struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *MjStoreItemsTopVo `json:"data,omitempty" xml:"data,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// total
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosStoreGetdefautitemsResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosStoreGetdefautitemsResultDo)
	},
}

// GetAlibabaMosStoreGetdefautitemsResultDo() 从对象池中获取AlibabaMosStoreGetdefautitemsResultDo
func GetAlibabaMosStoreGetdefautitemsResultDo() *AlibabaMosStoreGetdefautitemsResultDo {
	return poolAlibabaMosStoreGetdefautitemsResultDo.Get().(*AlibabaMosStoreGetdefautitemsResultDo)
}

// ReleaseAlibabaMosStoreGetdefautitemsResultDo 释放AlibabaMosStoreGetdefautitemsResultDo
func ReleaseAlibabaMosStoreGetdefautitemsResultDo(v *AlibabaMosStoreGetdefautitemsResultDo) {
	v.ErrMsg = ""
	v.Data = nil
	v.ErrCode = 0
	v.Total = 0
	v.Success = false
	poolAlibabaMosStoreGetdefautitemsResultDo.Put(v)
}
