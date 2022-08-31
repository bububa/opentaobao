package btrip

// OpenCostCenterAddEntityRs 结构体
type OpenCostCenterAddEntityRs struct {
	// 该成本中心下员工总数
	SelectedUserNum int64 `json:"selected_user_num,omitempty" xml:"selected_user_num,omitempty"`
	// 增加的人员信息条数
	AddNum int64 `json:"add_num,omitempty" xml:"add_num,omitempty"`
}
