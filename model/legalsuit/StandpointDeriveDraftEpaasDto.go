package legalsuit

// StandpointDeriveDraftEpaasDto 结构体
type StandpointDeriveDraftEpaasDto struct {
	// 业务id
	BusId string `json:"bus_id,omitempty" xml:"bus_id,omitempty"`
	// 来源
	DraftSource string `json:"draft_source,omitempty" xml:"draft_source,omitempty"`
	// 拓展字段json
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 场景名称
	SceneName string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// 业务名称
	BusName string `json:"bus_name,omitempty" xml:"bus_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 系统标识
	InputSystemCode string `json:"input_system_code,omitempty" xml:"input_system_code,omitempty"`
	// 编号
	CaseNo string `json:"case_no,omitempty" xml:"case_no,omitempty"`
	// 类型
	CaseType string `json:"case_type,omitempty" xml:"case_type,omitempty"`
	// 提交人
	SubmitPeople string `json:"submit_people,omitempty" xml:"submit_people,omitempty"`
	// 口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// 来源口径id
	SourceStandpointIds string `json:"source_standpoint_ids,omitempty" xml:"source_standpoint_ids,omitempty"`
	// 口径描述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
}
