package cainiaohandover

import (
	"sync"
)

// Courierlist 结构体
type Courierlist struct {
	// 承运商名字
	CourierName string `json:"courier_name,omitempty" xml:"courier_name,omitempty"`
	// 承运商code
	CourierCode string `json:"courier_code,omitempty" xml:"courier_code,omitempty"`
}

var poolCourierlist = sync.Pool{
	New: func() any {
		return new(Courierlist)
	},
}

// GetCourierlist() 从对象池中获取Courierlist
func GetCourierlist() *Courierlist {
	return poolCourierlist.Get().(*Courierlist)
}

// ReleaseCourierlist 释放Courierlist
func ReleaseCourierlist(v *Courierlist) {
	v.CourierName = ""
	v.CourierCode = ""
	poolCourierlist.Put(v)
}
