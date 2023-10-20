package xhotelonlineorder

import (
	"sync"
)

// CheckDateDo 结构体
type CheckDateDo struct {
	// 入住时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 离店时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
}

var poolCheckDateDo = sync.Pool{
	New: func() any {
		return new(CheckDateDo)
	},
}

// GetCheckDateDo() 从对象池中获取CheckDateDo
func GetCheckDateDo() *CheckDateDo {
	return poolCheckDateDo.Get().(*CheckDateDo)
}

// ReleaseCheckDateDo 释放CheckDateDo
func ReleaseCheckDateDo(v *CheckDateDo) {
	v.CheckOut = ""
	v.CheckIn = ""
	poolCheckDateDo.Put(v)
}
