package btrip

import (
	"sync"
)

// OpenCostCenterSaveRs 结构体
type OpenCostCenterSaveRs struct {
	// 商旅成本中心id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolOpenCostCenterSaveRs = sync.Pool{
	New: func() any {
		return new(OpenCostCenterSaveRs)
	},
}

// GetOpenCostCenterSaveRs() 从对象池中获取OpenCostCenterSaveRs
func GetOpenCostCenterSaveRs() *OpenCostCenterSaveRs {
	return poolOpenCostCenterSaveRs.Get().(*OpenCostCenterSaveRs)
}

// ReleaseOpenCostCenterSaveRs 释放OpenCostCenterSaveRs
func ReleaseOpenCostCenterSaveRs(v *OpenCostCenterSaveRs) {
	v.Id = 0
	poolOpenCostCenterSaveRs.Put(v)
}
