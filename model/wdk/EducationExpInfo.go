package wdk

import (
	"sync"
)

// EducationExpInfo 结构体
type EducationExpInfo struct {
	// 学历
	Education string `json:"education,omitempty" xml:"education,omitempty"`
	// 开始日期
	GmtEnd string `json:"gmt_end,omitempty" xml:"gmt_end,omitempty"`
	// 结束日期
	GmtStart string `json:"gmt_start,omitempty" xml:"gmt_start,omitempty"`
	// 专业
	Major string `json:"major,omitempty" xml:"major,omitempty"`
	// 学校
	School string `json:"school,omitempty" xml:"school,omitempty"`
	// 学制年数
	SchoolingYears string `json:"schooling_years,omitempty" xml:"schooling_years,omitempty"`
}

var poolEducationExpInfo = sync.Pool{
	New: func() any {
		return new(EducationExpInfo)
	},
}

// GetEducationExpInfo() 从对象池中获取EducationExpInfo
func GetEducationExpInfo() *EducationExpInfo {
	return poolEducationExpInfo.Get().(*EducationExpInfo)
}

// ReleaseEducationExpInfo 释放EducationExpInfo
func ReleaseEducationExpInfo(v *EducationExpInfo) {
	v.Education = ""
	v.GmtEnd = ""
	v.GmtStart = ""
	v.Major = ""
	v.School = ""
	v.SchoolingYears = ""
	poolEducationExpInfo.Put(v)
}
