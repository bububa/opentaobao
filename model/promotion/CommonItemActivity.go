package promotion

import (
	"sync"
)

// CommonItemActivity 结构体
type CommonItemActivity struct {
	// 活动描述，不能超过100字符
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 活动名称，不能超过32字符
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 人群标签值
	UserTag string `json:"user_tag,omitempty" xml:"user_tag,omitempty"`
	// 优惠活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 是否指定人群标签
	IsUserTag bool `json:"is_user_tag,omitempty" xml:"is_user_tag,omitempty"`
}

var poolCommonItemActivity = sync.Pool{
	New: func() any {
		return new(CommonItemActivity)
	},
}

// GetCommonItemActivity() 从对象池中获取CommonItemActivity
func GetCommonItemActivity() *CommonItemActivity {
	return poolCommonItemActivity.Get().(*CommonItemActivity)
}

// ReleaseCommonItemActivity 释放CommonItemActivity
func ReleaseCommonItemActivity(v *CommonItemActivity) {
	v.Description = ""
	v.EndTime = ""
	v.Name = ""
	v.StartTime = ""
	v.UserTag = ""
	v.ActivityId = 0
	v.IsUserTag = false
	poolCommonItemActivity.Put(v)
}
