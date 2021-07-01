package btrip

// OpenCostCenterSetEntityRs 结构体
type OpenCostCenterSetEntityRs struct {
	// 增加的人员信息条数
	AddNum int64 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	// 删除的人员信息条数
	RemoveNum int64 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	// 该成本中心下员工总数
	SelectedUserNum int64 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
}
