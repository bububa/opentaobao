package drugtrace

import (
	"sync"
)

// Attrinfolist 结构体
type Attrinfolist struct {
	// 货品id
	TracUserProductInfoId string `json:"trac_user_product_info_id,omitempty" xml:"trac_user_product_info_id,omitempty"`
	// 属性名称
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
}

var poolAttrinfolist = sync.Pool{
	New: func() any {
		return new(Attrinfolist)
	},
}

// GetAttrinfolist() 从对象池中获取Attrinfolist
func GetAttrinfolist() *Attrinfolist {
	return poolAttrinfolist.Get().(*Attrinfolist)
}

// ReleaseAttrinfolist 释放Attrinfolist
func ReleaseAttrinfolist(v *Attrinfolist) {
	v.TracUserProductInfoId = ""
	v.AttrName = ""
	poolAttrinfolist.Put(v)
}
