package eleenterpriserestaurant

import (
	"sync"
)

// Activite 结构体
type Activite struct {
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 餐厅描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 见附录【活动信息detail_type值】
	DetailType int64 `json:"detail_type,omitempty" xml:"detail_type,omitempty"`
	// 活动ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolActivite = sync.Pool{
	New: func() any {
		return new(Activite)
	},
}

// GetActivite() 从对象池中获取Activite
func GetActivite() *Activite {
	return poolActivite.Get().(*Activite)
}

// ReleaseActivite 释放Activite
func ReleaseActivite(v *Activite) {
	v.Name = ""
	v.Description = ""
	v.DetailType = 0
	v.Id = 0
	poolActivite.Put(v)
}
