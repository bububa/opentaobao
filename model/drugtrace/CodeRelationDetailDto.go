package drugtrace

// CodeRelationDetailDTO 
type CodeRelationDetailDTO struct {
    // 关系详情列表
    CodeRelationDetailList   []CodeRelationDetailListDO `json:"code_relation_detail_list,omitempty" xml:"code_relation_detail_list>code_relation_detail_list_do,omitempty"`
    // 文件信息
    CodeRelationDetailInfo   *CodeRelationDetailInfoDO `json:"code_relation_detail_info,omitempty" xml:"code_relation_detail_info,omitempty"`
}
