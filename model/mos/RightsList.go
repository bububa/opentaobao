package mos

import (
	"sync"
)

// RightsList 结构体
type RightsList struct {
	// 券名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 券价值
	CouponAmount int64 `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
	// id
	SnapshotId int64 `json:"snapshot_id,omitempty" xml:"snapshot_id,omitempty"`
	// 券使用门槛
	EntryAmount int64 `json:"entry_amount,omitempty" xml:"entry_amount,omitempty"`
}

var poolRightsList = sync.Pool{
	New: func() any {
		return new(RightsList)
	},
}

// GetRightsList() 从对象池中获取RightsList
func GetRightsList() *RightsList {
	return poolRightsList.Get().(*RightsList)
}

// ReleaseRightsList 释放RightsList
func ReleaseRightsList(v *RightsList) {
	v.Name = ""
	v.StartTime = ""
	v.EndTime = ""
	v.CouponAmount = 0
	v.SnapshotId = 0
	v.EntryAmount = 0
	poolRightsList.Put(v)
}
