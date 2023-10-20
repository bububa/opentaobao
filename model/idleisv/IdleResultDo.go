package idleisv

import (
	"sync"
)

// IdleResultDo 结构体
type IdleResultDo struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回数据
	Data *IdleItemApiDo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolIdleResultDo = sync.Pool{
	New: func() any {
		return new(IdleResultDo)
	},
}

// GetIdleResultDo() 从对象池中获取IdleResultDo
func GetIdleResultDo() *IdleResultDo {
	return poolIdleResultDo.Get().(*IdleResultDo)
}

// ReleaseIdleResultDo 释放IdleResultDo
func ReleaseIdleResultDo(v *IdleResultDo) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	v.Success = false
	poolIdleResultDo.Put(v)
}
