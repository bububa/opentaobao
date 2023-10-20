package icbulogistics

import (
	"sync"
)

// Country 结构体
type Country struct {
	// 地址代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 地址名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 未知
	PhoneCode string `json:"phone_code,omitempty" xml:"phone_code,omitempty"`
	// 地址id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolCountry = sync.Pool{
	New: func() any {
		return new(Country)
	},
}

// GetCountry() 从对象池中获取Country
func GetCountry() *Country {
	return poolCountry.Get().(*Country)
}

// ReleaseCountry 释放Country
func ReleaseCountry(v *Country) {
	v.Code = ""
	v.Name = ""
	v.PhoneCode = ""
	v.AreaId = 0
	poolCountry.Put(v)
}
