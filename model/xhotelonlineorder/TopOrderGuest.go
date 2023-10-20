package xhotelonlineorder

import (
	"sync"
)

// TopOrderGuest 结构体
type TopOrderGuest struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 住客类型
	CustomerType int64 `json:"customer_type,omitempty" xml:"customer_type,omitempty"`
	// 住客编号
	PersonNo int64 `json:"person_no,omitempty" xml:"person_no,omitempty"`
	// 年龄
	Age int64 `json:"age,omitempty" xml:"age,omitempty"`
	// 房间编号
	RoomNo int64 `json:"room_no,omitempty" xml:"room_no,omitempty"`
}

var poolTopOrderGuest = sync.Pool{
	New: func() any {
		return new(TopOrderGuest)
	},
}

// GetTopOrderGuest() 从对象池中获取TopOrderGuest
func GetTopOrderGuest() *TopOrderGuest {
	return poolTopOrderGuest.Get().(*TopOrderGuest)
}

// ReleaseTopOrderGuest 释放TopOrderGuest
func ReleaseTopOrderGuest(v *TopOrderGuest) {
	v.Name = ""
	v.CustomerType = 0
	v.PersonNo = 0
	v.Age = 0
	v.RoomNo = 0
	poolTopOrderGuest.Put(v)
}
