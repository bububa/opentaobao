package icbudropshipping

import (
	"sync"
)

// Phone 结构体
type Phone struct {
	// fax area
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// fax country
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// fax number
	Number string `json:"number,omitempty" xml:"number,omitempty"`
}

var poolPhone = sync.Pool{
	New: func() any {
		return new(Phone)
	},
}

// GetPhone() 从对象池中获取Phone
func GetPhone() *Phone {
	return poolPhone.Get().(*Phone)
}

// ReleasePhone 释放Phone
func ReleasePhone(v *Phone) {
	v.Area = ""
	v.Country = ""
	v.Number = ""
	poolPhone.Put(v)
}
