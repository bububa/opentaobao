package btrip

import (
	"sync"
)

// OpenCostCenterSetEntityRs 结构体
type OpenCostCenterSetEntityRs struct {
	// 增加的人员信息条数
	AddNum int64 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	// 删除的人员信息条数
	RemoveNum int64 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	// 该成本中心下员工总数
	SelectedUserNum int64 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}

var poolOpenCostCenterSetEntityRs = sync.Pool{
	New: func() any {
		return new(OpenCostCenterSetEntityRs)
	},
}

// GetOpenCostCenterSetEntityRs() 从对象池中获取OpenCostCenterSetEntityRs
func GetOpenCostCenterSetEntityRs() *OpenCostCenterSetEntityRs {
	return poolOpenCostCenterSetEntityRs.Get().(*OpenCostCenterSetEntityRs)
}

// ReleaseOpenCostCenterSetEntityRs 释放OpenCostCenterSetEntityRs
func ReleaseOpenCostCenterSetEntityRs(v *OpenCostCenterSetEntityRs) {
	v.AddNum = 0
	v.RemoveNum = 0
	v.SelectedUserNum = 0
	poolOpenCostCenterSetEntityRs.Put(v)
}
