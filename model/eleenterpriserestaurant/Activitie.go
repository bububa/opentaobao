package eleenterpriserestaurant

import (
	"sync"
)

// Activitie 结构体
type Activitie struct {
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 餐厅描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 见附录【活动信息detail_type值】
	DetailType int64 `json:"detail_type,omitempty" xml:"detail_type,omitempty"`
	// 活动ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolActivitie = sync.Pool{
	New: func() any {
		return new(Activitie)
	},
}

// GetActivitie() 从对象池中获取Activitie
func GetActivitie() *Activitie {
	return poolActivitie.Get().(*Activitie)
}

// ReleaseActivitie 释放Activitie
func ReleaseActivitie(v *Activitie) {
	v.Name = ""
	v.Description = ""
	v.DetailType = 0
	v.Id = 0
	poolActivitie.Put(v)
}
