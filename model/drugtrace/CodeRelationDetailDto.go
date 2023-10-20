package drugtrace

import (
	"sync"
)

// CodeRelationDetailDto 结构体
type CodeRelationDetailDto struct {
	// 关系详情列表
	CodeRelationDetailList []CodeRelationDetailListDo `json:"code_relation_detail_list,omitempty" xml:"code_relation_detail_list>code_relation_detail_list_do,omitempty"`
	// 文件信息
	CodeRelationDetailInfo *CodeRelationDetailInfoDo `json:"code_relation_detail_info,omitempty" xml:"code_relation_detail_info,omitempty"`
}

var poolCodeRelationDetailDto = sync.Pool{
	New: func() any {
		return new(CodeRelationDetailDto)
	},
}

// GetCodeRelationDetailDto() 从对象池中获取CodeRelationDetailDto
func GetCodeRelationDetailDto() *CodeRelationDetailDto {
	return poolCodeRelationDetailDto.Get().(*CodeRelationDetailDto)
}

// ReleaseCodeRelationDetailDto 释放CodeRelationDetailDto
func ReleaseCodeRelationDetailDto(v *CodeRelationDetailDto) {
	v.CodeRelationDetailList = v.CodeRelationDetailList[:0]
	v.CodeRelationDetailInfo = nil
	poolCodeRelationDetailDto.Put(v)
}
