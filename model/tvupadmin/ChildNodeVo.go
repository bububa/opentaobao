package tvupadmin

import (
	"sync"
)

// ChildNodeVo 结构体
type ChildNodeVo struct {
	// 显示名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 父类目名称
	ParentName string `json:"parent_name,omitempty" xml:"parent_name,omitempty"`
	// 主键ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 父类目ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
}

var poolChildNodeVo = sync.Pool{
	New: func() any {
		return new(ChildNodeVo)
	},
}

// GetChildNodeVo() 从对象池中获取ChildNodeVo
func GetChildNodeVo() *ChildNodeVo {
	return poolChildNodeVo.Get().(*ChildNodeVo)
}

// ReleaseChildNodeVo 释放ChildNodeVo
func ReleaseChildNodeVo(v *ChildNodeVo) {
	v.Name = ""
	v.ParentName = ""
	v.Id = 0
	v.ParentId = 0
	poolChildNodeVo.Put(v)
}
