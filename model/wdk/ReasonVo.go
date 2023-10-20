package wdk

import (
	"sync"
)

// ReasonVo 结构体
type ReasonVo struct {
	// 标签
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 原因描述
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// tip
	ReasonTip string `json:"reason_tip,omitempty" xml:"reason_tip,omitempty"`
	// 原因id
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}

var poolReasonVo = sync.Pool{
	New: func() any {
		return new(ReasonVo)
	},
}

// GetReasonVo() 从对象池中获取ReasonVo
func GetReasonVo() *ReasonVo {
	return poolReasonVo.Get().(*ReasonVo)
}

// ReleaseReasonVo 释放ReasonVo
func ReleaseReasonVo(v *ReasonVo) {
	v.Tags = v.Tags[:0]
	v.ReasonText = ""
	v.ReasonTip = ""
	v.ReasonId = 0
	poolReasonVo.Put(v)
}
