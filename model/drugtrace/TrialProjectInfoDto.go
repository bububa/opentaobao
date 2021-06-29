package drugtrace

// TrialProjectInfoDto 
type TrialProjectInfoDto struct {
    // 项目和药物属性
    ProjectAndAttributeDtoList   []TrialProjectAndAttributeDto `json:"project_and_attribute_dto_list,omitempty" xml:"project_and_attribute_dto_list>trial_project_and_attribute_dto,omitempty"`
    // 药品信息
    DrugDtoList   []SubType `json:"drug_dto_list,omitempty" xml:"drug_dto_list>sub_type,omitempty"`
}
