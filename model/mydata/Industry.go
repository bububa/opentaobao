package mydata

import (
	"sync"
)

// Industry 结构体
type Industry struct {
	// 行业描述
	IndustryDesc string `json:"industry_desc,omitempty" xml:"industry_desc,omitempty"`
	// 行业ID
	IndustryId int64 `json:"industry_id,omitempty" xml:"industry_id,omitempty"`
	// 是否主营行业
	MainCategory bool `json:"main_category,omitempty" xml:"main_category,omitempty"`
}

var poolIndustry = sync.Pool{
	New: func() any {
		return new(Industry)
	},
}

// GetIndustry() 从对象池中获取Industry
func GetIndustry() *Industry {
	return poolIndustry.Get().(*Industry)
}

// ReleaseIndustry 释放Industry
func ReleaseIndustry(v *Industry) {
	v.IndustryDesc = ""
	v.IndustryId = 0
	v.MainCategory = false
	poolIndustry.Put(v)
}
