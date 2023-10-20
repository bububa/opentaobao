package btrip

import (
	"sync"
)

// TgqNodeDo 结构体
type TgqNodeDo struct {
	// 规则列表
	DetailList []TimeNodeDo `json:"detail_list,omitempty" xml:"detail_list>time_node_do,omitempty"`
	// 是否支持
	Able bool `json:"able,omitempty" xml:"able,omitempty"`
}

var poolTgqNodeDo = sync.Pool{
	New: func() any {
		return new(TgqNodeDo)
	},
}

// GetTgqNodeDo() 从对象池中获取TgqNodeDo
func GetTgqNodeDo() *TgqNodeDo {
	return poolTgqNodeDo.Get().(*TgqNodeDo)
}

// ReleaseTgqNodeDo 释放TgqNodeDo
func ReleaseTgqNodeDo(v *TgqNodeDo) {
	v.DetailList = v.DetailList[:0]
	v.Able = false
	poolTgqNodeDo.Put(v)
}
