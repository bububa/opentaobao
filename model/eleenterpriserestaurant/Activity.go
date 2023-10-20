package eleenterpriserestaurant

import (
	"sync"
)

// Activity 结构体
type Activity struct {
	// 活动名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}

var poolActivity = sync.Pool{
	New: func() any {
		return new(Activity)
	},
}

// GetActivity() 从对象池中获取Activity
func GetActivity() *Activity {
	return poolActivity.Get().(*Activity)
}

// ReleaseActivity 释放Activity
func ReleaseActivity(v *Activity) {
	v.Name = ""
	v.Description = ""
	poolActivity.Put(v)
}
