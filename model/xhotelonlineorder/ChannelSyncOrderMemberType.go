package xhotelonlineorder

import (
	"sync"
)

// ChannelSyncOrderMemberType 结构体
type ChannelSyncOrderMemberType struct {
	// 会员类型枚举
	MemberType string `json:"member_type,omitempty" xml:"member_type,omitempty"`
	// 订单TID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolChannelSyncOrderMemberType = sync.Pool{
	New: func() any {
		return new(ChannelSyncOrderMemberType)
	},
}

// GetChannelSyncOrderMemberType() 从对象池中获取ChannelSyncOrderMemberType
func GetChannelSyncOrderMemberType() *ChannelSyncOrderMemberType {
	return poolChannelSyncOrderMemberType.Get().(*ChannelSyncOrderMemberType)
}

// ReleaseChannelSyncOrderMemberType 释放ChannelSyncOrderMemberType
func ReleaseChannelSyncOrderMemberType(v *ChannelSyncOrderMemberType) {
	v.MemberType = ""
	v.Tid = 0
	poolChannelSyncOrderMemberType.Put(v)
}
