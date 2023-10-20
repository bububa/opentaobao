package alitripmerchant

import (
	"sync"
)

// MemberRightVo 结构体
type MemberRightVo struct {
	// 权益名称
	MemberRightDesc string `json:"member_right_desc,omitempty" xml:"member_right_desc,omitempty"`
	// 权益内容
	MemberRightContent string `json:"member_right_content,omitempty" xml:"member_right_content,omitempty"`
	// 权益码点
	IconCodePoint string `json:"icon_code_point,omitempty" xml:"icon_code_point,omitempty"`
	// 顺序
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
}

var poolMemberRightVo = sync.Pool{
	New: func() any {
		return new(MemberRightVo)
	},
}

// GetMemberRightVo() 从对象池中获取MemberRightVo
func GetMemberRightVo() *MemberRightVo {
	return poolMemberRightVo.Get().(*MemberRightVo)
}

// ReleaseMemberRightVo 释放MemberRightVo
func ReleaseMemberRightVo(v *MemberRightVo) {
	v.MemberRightDesc = ""
	v.MemberRightContent = ""
	v.IconCodePoint = ""
	v.Order = 0
	poolMemberRightVo.Put(v)
}
