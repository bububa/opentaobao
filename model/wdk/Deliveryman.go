package wdk

import (
	"sync"
)

// Deliveryman 结构体
type Deliveryman struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 编号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolDeliveryman = sync.Pool{
	New: func() any {
		return new(Deliveryman)
	},
}

// GetDeliveryman() 从对象池中获取Deliveryman
func GetDeliveryman() *Deliveryman {
	return poolDeliveryman.Get().(*Deliveryman)
}

// ReleaseDeliveryman 释放Deliveryman
func ReleaseDeliveryman(v *Deliveryman) {
	v.Name = ""
	v.Code = ""
	v.Phone = ""
	poolDeliveryman.Put(v)
}
