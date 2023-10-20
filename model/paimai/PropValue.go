package paimai

import (
	"sync"
)

// PropValue 结构体
type PropValue struct {
	// 属性值
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 属性名
	PropName string `json:"prop_name,omitempty" xml:"prop_name,omitempty"`
	// 状态。可选值:normal(正常),deleted(删除)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 属性值别名
	NameAlias string `json:"name_alias,omitempty" xml:"name_alias,omitempty"`
	// 类目ID
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
	// 属性 ID
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// 属性值ID
	Vid int64 `json:"vid,omitempty" xml:"vid,omitempty"`
	// 排列序号。取值范围:大于零的整数
	SortOrder int64 `json:"sort_order,omitempty" xml:"sort_order,omitempty"`
	// 是否为父类目属性
	IsParent bool `json:"is_parent,omitempty" xml:"is_parent,omitempty"`
}

var poolPropValue = sync.Pool{
	New: func() any {
		return new(PropValue)
	},
}

// GetPropValue() 从对象池中获取PropValue
func GetPropValue() *PropValue {
	return poolPropValue.Get().(*PropValue)
}

// ReleasePropValue 释放PropValue
func ReleasePropValue(v *PropValue) {
	v.Name = ""
	v.PropName = ""
	v.Status = ""
	v.NameAlias = ""
	v.Cid = 0
	v.Pid = 0
	v.Vid = 0
	v.SortOrder = 0
	v.IsParent = false
	poolPropValue.Put(v)
}
