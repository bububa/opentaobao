package icbulogistics

import (
	"sync"
)

// Division 结构体
type Division struct {
	// 节点名称拼音
	Pinyin string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
	// 中文名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 上级节点ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 层级
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}

var poolDivision = sync.Pool{
	New: func() any {
		return new(Division)
	},
}

// GetDivision() 从对象池中获取Division
func GetDivision() *Division {
	return poolDivision.Get().(*Division)
}

// ReleaseDivision 释放Division
func ReleaseDivision(v *Division) {
	v.Pinyin = ""
	v.Name = ""
	v.Id = 0
	v.ParentId = 0
	v.Level = 0
	poolDivision.Put(v)
}
