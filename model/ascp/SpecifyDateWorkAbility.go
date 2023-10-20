package ascp

import (
	"sync"
)

// SpecifyDateWorkAbility 结构体
type SpecifyDateWorkAbility struct {
	// 指定日期，YYYY-MM-DD
	SpecifyDate string `json:"specify_date,omitempty" xml:"specify_date,omitempty"`
	// 上班时间(HH:MM:SS)
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 下班时间(HH:MM:SS)
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 即时单上门能力,举例 1 - 1小时内上门 2 - 2小时内上门 3 - 3小时内上门 4 - 4小时内上门 5 - 不支持即时单 未填写时，默认可2小时内上门
	ImmediateCollectAbility int64 `json:"immediate_collect_ability,omitempty" xml:"immediate_collect_ability,omitempty"`
	// 预约单上门能力，枚举： 1 - 1小时预约单可上门 2 - 2小时预约单可上门 3 - 半天预约单（上下午）可上门 4 - 当天预约单可上门
	ReservationAbility int64 `json:"reservation_ability,omitempty" xml:"reservation_ability,omitempty"`
}

var poolSpecifyDateWorkAbility = sync.Pool{
	New: func() any {
		return new(SpecifyDateWorkAbility)
	},
}

// GetSpecifyDateWorkAbility() 从对象池中获取SpecifyDateWorkAbility
func GetSpecifyDateWorkAbility() *SpecifyDateWorkAbility {
	return poolSpecifyDateWorkAbility.Get().(*SpecifyDateWorkAbility)
}

// ReleaseSpecifyDateWorkAbility 释放SpecifyDateWorkAbility
func ReleaseSpecifyDateWorkAbility(v *SpecifyDateWorkAbility) {
	v.SpecifyDate = ""
	v.BeginTime = ""
	v.EndTime = ""
	v.ImmediateCollectAbility = 0
	v.ReservationAbility = 0
	poolSpecifyDateWorkAbility.Put(v)
}
