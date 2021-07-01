package drugtrace

// CodeRelationDetailDto 结构体
type CodeRelationDetailDto struct {
	// 关系详情列表
	CodeRelationDetailList []CodeRelationDetailListDo `json:"code_relation_detail_list,omitempty" xml:"code_relation_detail_list>code_relation_detail_list_do,omitempty"`
	// 文件信息
	CodeRelationDetailInfo *CodeRelationDetailInfoDo `json:"code_relation_detail_info,omitempty" xml:"code_relation_detail_info,omitempty"`
}
