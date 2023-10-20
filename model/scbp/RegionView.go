package scbp

import (
	"sync"
)

// RegionView 结构体
type RegionView struct {
	// 国家列表
	CountryList []CountryView `json:"country_list,omitempty" xml:"country_list>country_view,omitempty"`
	// 地区中文名
	RegionCnName string `json:"region_cn_name,omitempty" xml:"region_cn_name,omitempty"`
}

var poolRegionView = sync.Pool{
	New: func() any {
		return new(RegionView)
	},
}

// GetRegionView() 从对象池中获取RegionView
func GetRegionView() *RegionView {
	return poolRegionView.Get().(*RegionView)
}

// ReleaseRegionView 释放RegionView
func ReleaseRegionView(v *RegionView) {
	v.CountryList = v.CountryList[:0]
	v.RegionCnName = ""
	poolRegionView.Put(v)
}
