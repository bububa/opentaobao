package btrip

import (
	"sync"
)

// OpenCostCenterDeleteEntityRs 结构体
type OpenCostCenterDeleteEntityRs struct {
	// 该成本中心下员工总数
	SelectedUserNum int64 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
	// 增加的人员信息条数
	RemoveNum int64 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
}

var poolOpenCostCenterDeleteEntityRs = sync.Pool{
	New: func() any {
		return new(OpenCostCenterDeleteEntityRs)
	},
}

// GetOpenCostCenterDeleteEntityRs() 从对象池中获取OpenCostCenterDeleteEntityRs
func GetOpenCostCenterDeleteEntityRs() *OpenCostCenterDeleteEntityRs {
	return poolOpenCostCenterDeleteEntityRs.Get().(*OpenCostCenterDeleteEntityRs)
}

// ReleaseOpenCostCenterDeleteEntityRs 释放OpenCostCenterDeleteEntityRs
func ReleaseOpenCostCenterDeleteEntityRs(v *OpenCostCenterDeleteEntityRs) {
	v.SelectedUserNum = 0
	v.RemoveNum = 0
	poolOpenCostCenterDeleteEntityRs.Put(v)
}
