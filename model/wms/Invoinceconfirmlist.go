package wms

import (
	"sync"
)

// Invoinceconfirmlist 结构体
type Invoinceconfirmlist struct {
	// 发票确认信息
	InvoinceConfirm *Invoinceconfirm `json:"invoince_confirm,omitempty" xml:"invoince_confirm,omitempty"`
}

var poolInvoinceconfirmlist = sync.Pool{
	New: func() any {
		return new(Invoinceconfirmlist)
	},
}

// GetInvoinceconfirmlist() 从对象池中获取Invoinceconfirmlist
func GetInvoinceconfirmlist() *Invoinceconfirmlist {
	return poolInvoinceconfirmlist.Get().(*Invoinceconfirmlist)
}

// ReleaseInvoinceconfirmlist 释放Invoinceconfirmlist
func ReleaseInvoinceconfirmlist(v *Invoinceconfirmlist) {
	v.InvoinceConfirm = nil
	poolInvoinceconfirmlist.Put(v)
}
