package btrip

import (
	"sync"
)

// OpenCostCenterAddEntityRs 结构体
type OpenCostCenterAddEntityRs struct {
	// 该成本中心下员工总数
	SelectedUserNum int64 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
	// 增加的人员信息条数
	AddNum int64 `json:"add_num,omitempty" xml:"add_num,omitempty"`
}

var poolOpenCostCenterAddEntityRs = sync.Pool{
	New: func() any {
		return new(OpenCostCenterAddEntityRs)
	},
}

// GetOpenCostCenterAddEntityRs() 从对象池中获取OpenCostCenterAddEntityRs
func GetOpenCostCenterAddEntityRs() *OpenCostCenterAddEntityRs {
	return poolOpenCostCenterAddEntityRs.Get().(*OpenCostCenterAddEntityRs)
}

// ReleaseOpenCostCenterAddEntityRs 释放OpenCostCenterAddEntityRs
func ReleaseOpenCostCenterAddEntityRs(v *OpenCostCenterAddEntityRs) {
	v.SelectedUserNum = 0
	v.AddNum = 0
	poolOpenCostCenterAddEntityRs.Put(v)
}
