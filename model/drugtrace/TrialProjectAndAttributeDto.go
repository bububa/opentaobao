package drugtrace

// TrialProjectAndAttributeDTO 
type TrialProjectAndAttributeDTO struct {
    // 项目名称
    TrialProjectName   string `json:"trial_project_name,omitempty" xml:"trial_project_name,omitempty"`
    // 项目编号
    TrialProjectNo   string `json:"trial_project_no,omitempty" xml:"trial_project_no,omitempty"`
    // 一级药物属性
    FirstAttrDtoList   []FirstAttrDTO `json:"first_attr_dto_list,omitempty" xml:"first_attr_dto_list>first_attr_dto,omitempty"`
}
