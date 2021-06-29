package drugtrace

// CodeRelationDto 
type CodeRelationDto struct {
    // 激活信息
    CodeActiveInfoDto   *CodeActiveInfoDto `json:"code_active_info_dto,omitempty" xml:"code_active_info_dto,omitempty"`
    // 码关联关系
    CodeRelationList   []CodeInfo `json:"code_relation_list,omitempty" xml:"code_relation_list>code_info,omitempty"`
    // 是否是最小包装
    IsSmallest   string `json:"is_smallest,omitempty" xml:"is_smallest,omitempty"`
    // 药品包装信息
    PkgInfoDto   *PkgInfoDto `json:"pkg_info_dto,omitempty" xml:"pkg_info_dto,omitempty"`
    // 药品基础信息
    BaseInfosDto   *BaseInfosDto `json:"base_infos_dto,omitempty" xml:"base_infos_dto,omitempty"`
    // 生产信息
    ProduceInfoList   []ProduceInfoDto `json:"produce_info_list,omitempty" xml:"produce_info_list>produce_info_dto,omitempty"`
    // 激活信息
    CodeActiveInfoDTO   *CodeActiveInfoDto `json:"code_active_info_d_t_o,omitempty" xml:"code_active_info_d_t_o,omitempty"`
    // 药品包装信息
    PkgInfoDTO   *PkgInfoDto `json:"pkg_info_d_t_o,omitempty" xml:"pkg_info_d_t_o,omitempty"`
    // 药品基础信息
    BaseInfosDTO   *BaseInfosDto `json:"base_infos_d_t_o,omitempty" xml:"base_infos_d_t_o,omitempty"`
    // errorCodeContent
    ErrorCodeContent   string `json:"error_code_content,omitempty" xml:"error_code_content,omitempty"`
}
