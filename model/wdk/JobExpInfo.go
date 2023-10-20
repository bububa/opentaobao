package wdk

import (
	"sync"
)

// JobExpInfo 结构体
type JobExpInfo struct {
	// 部门
	Department string `json:"department,omitempty" xml:"department,omitempty"`
	// 离职原因
	DimissionReason string `json:"dimission_reason,omitempty" xml:"dimission_reason,omitempty"`
	// 结束日期
	GmtEnd string `json:"gmt_end,omitempty" xml:"gmt_end,omitempty"`
	// 开始日期
	GmtStart string `json:"gmt_start,omitempty" xml:"gmt_start,omitempty"`
	// 职位
	Position string `json:"position,omitempty" xml:"position,omitempty"`
	// 薪资（月）
	SalaryByMonth string `json:"salary_by_month,omitempty" xml:"salary_by_month,omitempty"`
	// 工作单位
	Company string `json:"company,omitempty" xml:"company,omitempty"`
}

var poolJobExpInfo = sync.Pool{
	New: func() any {
		return new(JobExpInfo)
	},
}

// GetJobExpInfo() 从对象池中获取JobExpInfo
func GetJobExpInfo() *JobExpInfo {
	return poolJobExpInfo.Get().(*JobExpInfo)
}

// ReleaseJobExpInfo 释放JobExpInfo
func ReleaseJobExpInfo(v *JobExpInfo) {
	v.Department = ""
	v.DimissionReason = ""
	v.GmtEnd = ""
	v.GmtStart = ""
	v.Position = ""
	v.SalaryByMonth = ""
	v.Company = ""
	poolJobExpInfo.Put(v)
}
