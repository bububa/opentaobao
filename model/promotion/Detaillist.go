package promotion

import (
	"sync"
)

// Detaillist 结构体
type Detaillist struct {
	// 需要消耗星星数
	UnitPoint string `json:"unit_point,omitempty" xml:"unit_point,omitempty"`
	// 扩展字段
	ExtMap string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 价格分
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}

var poolDetaillist = sync.Pool{
	New: func() any {
		return new(Detaillist)
	},
}

// GetDetaillist() 从对象池中获取Detaillist
func GetDetaillist() *Detaillist {
	return poolDetaillist.Get().(*Detaillist)
}

// ReleaseDetaillist 释放Detaillist
func ReleaseDetaillist(v *Detaillist) {
	v.UnitPoint = ""
	v.ExtMap = ""
	v.UnitPrice = 0
	v.Num = 0
	poolDetaillist.Put(v)
}
