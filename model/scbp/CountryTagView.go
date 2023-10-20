package scbp

import (
	"sync"
)

// CountryTagView 结构体
type CountryTagView struct {
	// 标签中文名
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 标签id
	TagId string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 最近7天效果数据
	Effect *Effect7d `json:"effect,omitempty" xml:"effect,omitempty"`
	// 溢价百分比
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolCountryTagView = sync.Pool{
	New: func() any {
		return new(CountryTagView)
	},
}

// GetCountryTagView() 从对象池中获取CountryTagView
func GetCountryTagView() *CountryTagView {
	return poolCountryTagView.Get().(*CountryTagView)
}

// ReleaseCountryTagView 释放CountryTagView
func ReleaseCountryTagView(v *CountryTagView) {
	v.TagName = ""
	v.TagId = ""
	v.Effect = nil
	v.Discount = 0
	poolCountryTagView.Put(v)
}
