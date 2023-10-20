package security

import (
	"sync"
)

// JaqDispatchParam 结构体
type JaqDispatchParam struct {
	// 事件ID
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 协议版本号
	ProtocolVersion string `json:"protocol_version,omitempty" xml:"protocol_version,omitempty"`
	// 下发的软token密文
	Rtken string `json:"rtken,omitempty" xml:"rtken,omitempty"`
	// 下发的软token索引
	RtkenIndex string `json:"rtken_index,omitempty" xml:"rtken_index,omitempty"`
}

var poolJaqDispatchParam = sync.Pool{
	New: func() any {
		return new(JaqDispatchParam)
	},
}

// GetJaqDispatchParam() 从对象池中获取JaqDispatchParam
func GetJaqDispatchParam() *JaqDispatchParam {
	return poolJaqDispatchParam.Get().(*JaqDispatchParam)
}

// ReleaseJaqDispatchParam 释放JaqDispatchParam
func ReleaseJaqDispatchParam(v *JaqDispatchParam) {
	v.EventId = ""
	v.ProtocolVersion = ""
	v.Rtken = ""
	v.RtkenIndex = ""
	poolJaqDispatchParam.Put(v)
}
