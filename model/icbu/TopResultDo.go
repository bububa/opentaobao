package icbu

import (
	"sync"
)

// TopResultDo 结构体
type TopResultDo struct {
	// 层级属性的下一级属性结构
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 用于排查系统错误
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
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
	v.BizSuccess = false
	poolTopResultDo.Put(v)
}
