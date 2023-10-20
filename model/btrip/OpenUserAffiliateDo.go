package btrip

import (
	"sync"
)

// OpenUserAffiliateDo 结构体
type OpenUserAffiliateDo struct {
	// 出行人ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 出行人名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

var poolOpenUserAffiliateDo = sync.Pool{
	New: func() any {
		return new(OpenUserAffiliateDo)
	},
}

// GetOpenUserAffiliateDo() 从对象池中获取OpenUserAffiliateDo
func GetOpenUserAffiliateDo() *OpenUserAffiliateDo {
	return poolOpenUserAffiliateDo.Get().(*OpenUserAffiliateDo)
}

// ReleaseOpenUserAffiliateDo 释放OpenUserAffiliateDo
func ReleaseOpenUserAffiliateDo(v *OpenUserAffiliateDo) {
	v.UserId = ""
	v.UserName = ""
	poolOpenUserAffiliateDo.Put(v)
}
