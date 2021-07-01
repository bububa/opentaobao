package drugtrace

// TrialProjectDto 结构体
type TrialProjectDto struct {
	// 项目名称
	TrialProjectName string `json:"trial_project_name,omitempty" xml:"trial_project_name,omitempty"`
	// 项目编号
	TrialProjectNo string `json:"trial_project_no,omitempty" xml:"trial_project_no,omitempty"`
	// 药品信息
	DrugDtoList []SubType `json:"drug_dto_list,omitempty" xml:"drug_dto_list>sub_type,omitempty"`
}
