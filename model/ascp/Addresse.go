package ascp

import (
	"sync"
)

// Addresse 结构体
type Addresse struct {
	// 收货地：省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 收货地：城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货地：区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 收货地：街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}

var poolAddresse = sync.Pool{
	New: func() any {
		return new(Addresse)
	},
}

// GetAddresse() 从对象池中获取Addresse
func GetAddresse() *Addresse {
	return poolAddresse.Get().(*Addresse)
}

// ReleaseAddresse 释放Addresse
func ReleaseAddresse(v *Addresse) {
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	poolAddresse.Put(v)
}
