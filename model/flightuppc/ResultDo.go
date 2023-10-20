package flightuppc

import (
	"sync"
)

// ResultDo 结构体
type ResultDo struct {
	// 错误信息
	MsgForClient string `json:"msg_for_client,omitempty" xml:"msg_for_client,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否回流成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
	// 是否执行回流成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultDo = sync.Pool{
	New: func() any {
		return new(ResultDo)
	},
}

// GetResultDo() 从对象池中获取ResultDo
func GetResultDo() *ResultDo {
	return poolResultDo.Get().(*ResultDo)
}

// ReleaseResultDo 释放ResultDo
func ReleaseResultDo(v *ResultDo) {
	v.MsgForClient = ""
	v.Code = ""
	v.Msg = ""
	v.Module = false
	v.Success = false
	poolResultDo.Put(v)
}
