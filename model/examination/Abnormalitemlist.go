package examination

import (
	"sync"
)

// Abnormalitemlist 结构体
type Abnormalitemlist struct {
	// 异常项指标名称
	ItemDetail string `json:"item_detail,omitempty" xml:"item_detail,omitempty"`
	// 异常项指标详情
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
}

var poolAbnormalitemlist = sync.Pool{
	New: func() any {
		return new(Abnormalitemlist)
	},
}

// GetAbnormalitemlist() 从对象池中获取Abnormalitemlist
func GetAbnormalitemlist() *Abnormalitemlist {
	return poolAbnormalitemlist.Get().(*Abnormalitemlist)
}

// ReleaseAbnormalitemlist 释放Abnormalitemlist
func ReleaseAbnormalitemlist(v *Abnormalitemlist) {
	v.ItemDetail = ""
	v.ItemName = ""
	poolAbnormalitemlist.Put(v)
}
