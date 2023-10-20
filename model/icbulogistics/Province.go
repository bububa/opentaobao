package icbulogistics

import (
	"sync"
)

// Province 结构体
type Province struct {
	// 地址代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 地址名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolProvince = sync.Pool{
	New: func() any {
		return new(Province)
	},
}

// GetProvince() 从对象池中获取Province
func GetProvince() *Province {
	return poolProvince.Get().(*Province)
}

// ReleaseProvince 释放Province
func ReleaseProvince(v *Province) {
	v.Code = ""
	v.Name = ""
	v.AreaId = 0
	poolProvince.Put(v)
}
