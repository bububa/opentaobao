package alihouse

// ProjectCooperationDto 结构体
type ProjectCooperationDto struct {
	// 合作开始时间
	CooperationEndTime string `json:"cooperation_end_time,omitempty" xml:"cooperation_end_time,omitempty"`
	// 合作结束时间
	CooperationBeginTime string `json:"cooperation_begin_time,omitempty" xml:"cooperation_begin_time,omitempty"`
	// KA楼盘ID
	KaProjectMid string `json:"ka_project_mid,omitempty" xml:"ka_project_mid,omitempty"`
	// 流水号
	OuterCooperationId string `json:"outer_cooperation_id,omitempty" xml:"outer_cooperation_id,omitempty"`
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// KA楼盘码
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 合同号
	ContractNo string `json:"contract_no,omitempty" xml:"contract_no,omitempty"`
	// 外部活动ID
	OuterActivityId string `json:"outer_activity_id,omitempty" xml:"outer_activity_id,omitempty"`
	// KA楼盘ID
	KaOuterId string `json:"ka_outer_id,omitempty" xml:"ka_outer_id,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 状态 0-无效 1-有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否优先展示 0-否 1-是
	IsPriority int64 `json:"is_priority,omitempty" xml:"is_priority,omitempty"`
	// 0-KA合作  10-优先合作
	CooperationType int64 `json:"cooperation_type,omitempty" xml:"cooperation_type,omitempty"`
	// 活动类型 1- 百亿房补大额电商券 2 - 直通车活动
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
}
