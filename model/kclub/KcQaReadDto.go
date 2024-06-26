package kclub

import (
	"sync"
)

// KcQaReadDto 结构体
type KcQaReadDto struct {
	// 子知识
	ChildQas []KcQaReadDto `json:"child_qas,omitempty" xml:"child_qas>kc_qa_read_dto,omitempty"`
	// 问题答案
	Solutions []KcQaSolutionDto `json:"solutions,omitempty" xml:"solutions>kc_qa_solution_dto,omitempty"`
	// 父类目id
	ParentCats []int64 `json:"parent_cats,omitempty" xml:"parent_cats>int64,omitempty"`
	// 类目路径
	CatPath string `json:"cat_path,omitempty" xml:"cat_path,omitempty"`
	// 知识标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 编辑时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 关联的实体code
	EntityCode string `json:"entity_code,omitempty" xml:"entity_code,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 问题id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 问题类型(原kbs context)
	Context int64 `json:"context,omitempty" xml:"context,omitempty"`
	// 问题类型
	QuestionType int64 `json:"question_type,omitempty" xml:"question_type,omitempty"`
	// qa的访问量
	QaPv *QaPvDto `json:"qa_pv,omitempty" xml:"qa_pv,omitempty"`
}

var poolKcQaReadDto = sync.Pool{
	New: func() any {
		return new(KcQaReadDto)
	},
}

// GetKcQaReadDto() 从对象池中获取KcQaReadDto
func GetKcQaReadDto() *KcQaReadDto {
	return poolKcQaReadDto.Get().(*KcQaReadDto)
}

// ReleaseKcQaReadDto 释放KcQaReadDto
func ReleaseKcQaReadDto(v *KcQaReadDto) {
	v.ChildQas = v.ChildQas[:0]
	v.Solutions = v.Solutions[:0]
	v.ParentCats = v.ParentCats[:0]
	v.CatPath = ""
	v.Title = ""
	v.GmtModified = ""
	v.GmtCreate = ""
	v.EntityCode = ""
	v.CatId = 0
	v.TenantId = 0
	v.Status = 0
	v.Id = 0
	v.Context = 0
	v.QuestionType = 0
	v.QaPv = nil
	poolKcQaReadDto.Put(v)
}
