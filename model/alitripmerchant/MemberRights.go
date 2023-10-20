package alitripmerchant

import (
	"sync"
)

// MemberRights 结构体
type MemberRights struct {
	// 权益图标
	MemberRightIcon string `json:"member_right_icon,omitempty" xml:"member_right_icon,omitempty"`
	// 权益描述
	MemberRightDesc string `json:"member_right_desc,omitempty" xml:"member_right_desc,omitempty"`
}

var poolMemberRights = sync.Pool{
	New: func() any {
		return new(MemberRights)
	},
}

// GetMemberRights() 从对象池中获取MemberRights
func GetMemberRights() *MemberRights {
	return poolMemberRights.Get().(*MemberRights)
}

// ReleaseMemberRights 释放MemberRights
func ReleaseMemberRights(v *MemberRights) {
	v.MemberRightIcon = ""
	v.MemberRightDesc = ""
	poolMemberRights.Put(v)
}
