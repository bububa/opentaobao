package scbp

import (
	"sync"
)

// CountryView 结构体
type CountryView struct {
	// 国家中文名
	CountryCnName string `json:"country_cn_name,omitempty" xml:"country_cn_name,omitempty"`
	// 国家ID
	CountryId string `json:"country_id,omitempty" xml:"country_id,omitempty"`
}

var poolCountryView = sync.Pool{
	New: func() any {
		return new(CountryView)
	},
}

// GetCountryView() 从对象池中获取CountryView
func GetCountryView() *CountryView {
	return poolCountryView.Get().(*CountryView)
}

// ReleaseCountryView 释放CountryView
func ReleaseCountryView(v *CountryView) {
	v.CountryCnName = ""
	v.CountryId = ""
	poolCountryView.Put(v)
}
