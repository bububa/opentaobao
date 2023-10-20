package ascp

import (
	"sync"
)

// Blacklist 结构体
type Blacklist struct {
	// 黑名单用户手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 黑名单用户类型（枚举），多个以英文逗号隔开 1-恶意下单 2-恶意投诉 3-恶意敲诈 4-人身攻击类（辱骂殴打快递员等） 5-其他
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
}

var poolBlacklist = sync.Pool{
	New: func() any {
		return new(Blacklist)
	},
}

// GetBlacklist() 从对象池中获取Blacklist
func GetBlacklist() *Blacklist {
	return poolBlacklist.Get().(*Blacklist)
}

// ReleaseBlacklist 释放Blacklist
func ReleaseBlacklist(v *Blacklist) {
	v.Mobile = ""
	v.Reason = ""
	poolBlacklist.Put(v)
}
