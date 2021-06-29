package drugtrace

// CodeQueryDrugInfoDTO 
type CodeQueryDrugInfoDTO struct {
    // 药品信息列表
    InfoList   []CodeFullInfoDTO `json:"info_list,omitempty" xml:"info_list>code_full_info_dto,omitempty"`
}
