package drugtrace

// CodeRelationDTO 
type CodeRelationDTO struct {
    // 激活信息
    CodeActiveInfoDto   *CodeActiveInfoDTO `json:"code_active_info_dto,omitempty" xml:"code_active_info_dto,omitempty"`
    // 码关联关系
    CodeRelationList   []CodeInfo `json:"code_relation_list,omitempty" xml:"code_relation_list>code_info,omitempty"`
    // 是否是最小包装
    IsSmallest   string `json:"is_smallest,omitempty" xml:"is_smallest,omitempty"`
    // 药品包装信息
    PkgInfoDto   *PkgInfoDTO `json:"pkg_info_dto,omitempty" xml:"pkg_info_dto,omitempty"`
    // 药品基础信息
    BaseInfosDto   *BaseInfosDTO `json:"base_infos_dto,omitempty" xml:"base_infos_dto,omitempty"`
    // 生产信息
    ProduceInfoList   []ProduceInfoDTO `json:"produce_info_list,omitempty" xml:"produce_info_list>produce_info_dto,omitempty"`
    // 激活信息
    CodeActiveInfoDTO   *CodeActiveInfoDTO `json:"code_active_info_d_t_o,omitempty" xml:"code_active_info_d_t_o,omitempty"`
    // 药品包装信息
    PkgInfoDTO   *PkgInfoDTO `json:"pkg_info_d_t_o,omitempty" xml:"pkg_info_d_t_o,omitempty"`
    // 药品基础信息
    BaseInfosDTO   *BaseInfosDTO `json:"base_infos_d_t_o,omitempty" xml:"base_infos_d_t_o,omitempty"`
    // 追溯码；查询的码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 父码
    ParentCode   string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
    // 码级别
    CodeLevel   string `json:"code_level,omitempty" xml:"code_level,omitempty"`
    // 包装级别
    CodePackLevel   string `json:"code_pack_level,omitempty" xml:"code_pack_level,omitempty"`
    // 装箱数量;小盒码，返回1；中包码，返回实际小盒数量；大箱码，返回实际小盒数量
    BoxAmount   int64 `json:"box_amount,omitempty" xml:"box_amount,omitempty"`
    // 大箱或中包状态；若扫描的是小盒码，直接返回正常； 0-正常；1-拼箱；2-零箱；3-即拼箱又零箱
    BoxStatus   int64 `json:"box_status,omitempty" xml:"box_status,omitempty"`
    // 码状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // errorCodeContent
    ErrorCodeContent   string `json:"error_code_content,omitempty" xml:"error_code_content,omitempty"`
}
