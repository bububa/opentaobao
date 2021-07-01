package drugtrace

// CodeQueryDrugInfoDto 
type CodeQueryDrugInfoDto struct {
    // 药品信息列表
    InfoList   []CodeFullInfoDto `json:"info_list,omitempty" xml:"info_list>code_full_info_dto,omitempty"`
}
