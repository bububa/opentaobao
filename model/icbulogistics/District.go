package icbulogistics

import (
	"sync"
)

// District 结构体
type District struct {
	// 地址代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 地址名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址代码
	AreaId string `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolDistrict = sync.Pool{
	New: func() any {
		return new(District)
	},
}

// GetDistrict() 从对象池中获取District
func GetDistrict() *District {
	return poolDistrict.Get().(*District)
}

// ReleaseDistrict 释放District
func ReleaseDistrict(v *District) {
	v.Code = ""
	v.Name = ""
	v.AreaId = ""
	poolDistrict.Put(v)
}
