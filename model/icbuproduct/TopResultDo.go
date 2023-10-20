package icbuproduct

import (
	"sync"
)

// TopResultDo 结构体
type TopResultDo struct {
	// 库存更新是否成功
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 接口更新失败时错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 失败错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口失败时的追踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopResultDo = sync.Pool{
	New: func() any {
		return new(TopResultDo)
	},
}

// GetTopResultDo() 从对象池中获取TopResultDo
func GetTopResultDo() *TopResultDo {
	return poolTopResultDo.Get().(*TopResultDo)
}

// ReleaseTopResultDo 释放TopResultDo
func ReleaseTopResultDo(v *TopResultDo) {
	v.Data = ""
	v.Message = ""
	v.MsgCode = ""
	v.TraceId = ""
	v.Success = false
	poolTopResultDo.Put(v)
}
