package promotion

import (
	"sync"
)

// FailElement 结构体
type FailElement struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 参与者id
	ParticipateId string `json:"participate_id,omitempty" xml:"participate_id,omitempty"`
	// 参与者名称
	ParticipateName string `json:"participate_name,omitempty" xml:"participate_name,omitempty"`
}

var poolFailElement = sync.Pool{
	New: func() any {
		return new(FailElement)
	},
}

// GetFailElement() 从对象池中获取FailElement
func GetFailElement() *FailElement {
	return poolFailElement.Get().(*FailElement)
}

// ReleaseFailElement 释放FailElement
func ReleaseFailElement(v *FailElement) {
	v.ErrorMsg = ""
	v.ParticipateId = ""
	v.ParticipateName = ""
	poolFailElement.Put(v)
}
