package btrip

import (
	"sync"
)

// Entity 结构体
type Entity struct {
	// 实体id，all_employe为false时，entities里元素的id必传
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 实体名，all_employe为false时，entities里元素的name必传
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 1：员工，all_employe为false时，entities里元素的type必传
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolEntity = sync.Pool{
	New: func() any {
		return new(Entity)
	},
}

// GetEntity() 从对象池中获取Entity
func GetEntity() *Entity {
	return poolEntity.Get().(*Entity)
}

// ReleaseEntity 释放Entity
func ReleaseEntity(v *Entity) {
	v.Id = ""
	v.Name = ""
	v.Type = 0
	poolEntity.Put(v)
}
