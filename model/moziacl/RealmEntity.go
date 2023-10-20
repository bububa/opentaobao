package moziacl

import (
	"sync"
)

// RealmEntity 结构体
type RealmEntity struct {
	// 租户名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 租户描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 租户id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolRealmEntity = sync.Pool{
	New: func() any {
		return new(RealmEntity)
	},
}

// GetRealmEntity() 从对象池中获取RealmEntity
func GetRealmEntity() *RealmEntity {
	return poolRealmEntity.Get().(*RealmEntity)
}

// ReleaseRealmEntity 释放RealmEntity
func ReleaseRealmEntity(v *RealmEntity) {
	v.Name = ""
	v.Description = ""
	v.Id = 0
	poolRealmEntity.Put(v)
}
