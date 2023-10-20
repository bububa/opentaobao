package jms

import (
	"sync"
)

// MsgGroupDo 结构体
type MsgGroupDo struct {
	// 123
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolMsgGroupDo = sync.Pool{
	New: func() any {
		return new(MsgGroupDo)
	},
}

// GetMsgGroupDo() 从对象池中获取MsgGroupDo
func GetMsgGroupDo() *MsgGroupDo {
	return poolMsgGroupDo.Get().(*MsgGroupDo)
}

// ReleaseMsgGroupDo 释放MsgGroupDo
func ReleaseMsgGroupDo(v *MsgGroupDo) {
	v.Name = ""
	poolMsgGroupDo.Put(v)
}
