package tvupadmin

import (
	"sync"
)

// UserRightsDo 结构体
type UserRightsDo struct {
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 权益截止时间
	GmtEnd string `json:"gmt_end,omitempty" xml:"gmt_end,omitempty"`
	// 权益开始时间
	GmtStart string `json:"gmt_start,omitempty" xml:"gmt_start,omitempty"`
	// 权益类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 权益ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 用户ID
	Uid int64 `json:"uid,omitempty" xml:"uid,omitempty"`
	// 是否连续包月
	RenewalSupport bool `json:"renewal_support,omitempty" xml:"renewal_support,omitempty"`
}

var poolUserRightsDo = sync.Pool{
	New: func() any {
		return new(UserRightsDo)
	},
}

// GetUserRightsDo() 从对象池中获取UserRightsDo
func GetUserRightsDo() *UserRightsDo {
	return poolUserRightsDo.Get().(*UserRightsDo)
}

// ReleaseUserRightsDo 释放UserRightsDo
func ReleaseUserRightsDo(v *UserRightsDo) {
	v.GmtModified = ""
	v.GmtCreate = ""
	v.GmtEnd = ""
	v.GmtStart = ""
	v.Type = ""
	v.ItemId = ""
	v.Uid = 0
	v.RenewalSupport = false
	poolUserRightsDo.Put(v)
}
