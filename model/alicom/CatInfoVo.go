package alicom

import (
	"sync"
)

// CatInfoVo 结构体
type CatInfoVo struct {
	// 手机归属区域
	ShowCatName string `json:"show_cat_name,omitempty" xml:"show_cat_name,omitempty"`
	// 手机归属运营商
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 手机归属城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 运营商
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
}

var poolCatInfoVo = sync.Pool{
	New: func() any {
		return new(CatInfoVo)
	},
}

// GetCatInfoVo() 从对象池中获取CatInfoVo
func GetCatInfoVo() *CatInfoVo {
	return poolCatInfoVo.Get().(*CatInfoVo)
}

// ReleaseCatInfoVo 释放CatInfoVo
func ReleaseCatInfoVo(v *CatInfoVo) {
	v.ShowCatName = ""
	v.AreaName = ""
	v.City = ""
	v.CatName = ""
	poolCatInfoVo.Put(v)
}
