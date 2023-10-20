package happytrip

import (
	"sync"
)

// PeerStaff 结构体
type PeerStaff struct {
	// 同行人用户id，阿里工号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 同行人姓名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

var poolPeerStaff = sync.Pool{
	New: func() any {
		return new(PeerStaff)
	},
}

// GetPeerStaff() 从对象池中获取PeerStaff
func GetPeerStaff() *PeerStaff {
	return poolPeerStaff.Get().(*PeerStaff)
}

// ReleasePeerStaff 释放PeerStaff
func ReleasePeerStaff(v *PeerStaff) {
	v.UserId = ""
	v.UserName = ""
	poolPeerStaff.Put(v)
}
