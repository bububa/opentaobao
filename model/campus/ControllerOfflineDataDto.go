package campus

// ControllerOfflineDataDto 结构体
type ControllerOfflineDataDto struct {
	// 门禁点数据
	GuardOfflineList []GuardOfflineDataDto `json:"guard_offline_list,omitempty" xml:"guard_offline_list>guard_offline_data_dto,omitempty"`
}
