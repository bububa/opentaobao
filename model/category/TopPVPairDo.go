package category

import (
	"sync"
)

// TopPVPairDo 结构体
type TopPVPairDo struct {
	// 属性项名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 属性值名称
	ValueName string `json:"value_name,omitempty" xml:"value_name,omitempty"`
	// 属性值ID
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 属性项ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}

var poolTopPVPairDo = sync.Pool{
	New: func() any {
		return new(TopPVPairDo)
	},
}

// GetTopPVPairDo() 从对象池中获取TopPVPairDo
func GetTopPVPairDo() *TopPVPairDo {
	return poolTopPVPairDo.Get().(*TopPVPairDo)
}

// ReleaseTopPVPairDo 释放TopPVPairDo
func ReleaseTopPVPairDo(v *TopPVPairDo) {
	v.PropertyName = ""
	v.ValueName = ""
	v.ValueId = 0
	v.PropertyId = 0
	poolTopPVPairDo.Put(v)
}
