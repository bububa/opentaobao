package simba

import (
	"sync"
)

// AreaOption 结构体
type AreaOption struct {
	// 地域名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地域id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 父地域id，若该字段为0表示该行政区为顶层，例如像北京，国外等。
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 地域级别，目前自治区、省、直辖市是1，其他城市、地区是2
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}

var poolAreaOption = sync.Pool{
	New: func() any {
		return new(AreaOption)
	},
}

// GetAreaOption() 从对象池中获取AreaOption
func GetAreaOption() *AreaOption {
	return poolAreaOption.Get().(*AreaOption)
}

// ReleaseAreaOption 释放AreaOption
func ReleaseAreaOption(v *AreaOption) {
	v.Name = ""
	v.AreaId = 0
	v.ParentId = 0
	v.Level = 0
	poolAreaOption.Put(v)
}
