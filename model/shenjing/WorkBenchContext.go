package shenjing

import (
	"sync"
)

// WorkBenchContext 结构体
type WorkBenchContext struct {
	// 授权的appCode
	AppCode string `json:"app_code,omitempty" xml:"app_code,omitempty"`
	// 系统ID
	SystemId string `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 园区Id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
}

var poolWorkBenchContext = sync.Pool{
	New: func() any {
		return new(WorkBenchContext)
	},
}

// GetWorkBenchContext() 从对象池中获取WorkBenchContext
func GetWorkBenchContext() *WorkBenchContext {
	return poolWorkBenchContext.Get().(*WorkBenchContext)
}

// ReleaseWorkBenchContext 释放WorkBenchContext
func ReleaseWorkBenchContext(v *WorkBenchContext) {
	v.AppCode = ""
	v.SystemId = ""
	v.CompanyId = 0
	v.CampusId = 0
	poolWorkBenchContext.Put(v)
}
