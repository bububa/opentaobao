package kclub

import (
	"sync"
)

// KcSearchQuestion 结构体
type KcSearchQuestion struct {
	// catId 路径
	CatIdPathList []int64 `json:"cat_id_path_list,omitempty" xml:"cat_id_path_list>int64,omitempty"`
	// 问题关联的实体code
	EntityCode string `json:"entity_code,omitempty" xml:"entity_code,omitempty"`
	// 编辑人名称
	ModifiedUserName string `json:"modified_user_name,omitempty" xml:"modified_user_name,omitempty"`
	// 创建人名称
	CreateUserName string `json:"create_user_name,omitempty" xml:"create_user_name,omitempty"`
	// 标题分词
	TitleSegment string `json:"title_segment,omitempty" xml:"title_segment,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 编辑时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// qa 失效时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// qa 生效时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 额外字段
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 类目路径
	CatPathName string `json:"cat_path_name,omitempty" xml:"cat_path_name,omitempty"`
	// 记录唯一编码
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 得分
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 相似标题
	SimilarTitle string `json:"similar_title,omitempty" xml:"similar_title,omitempty"`
	// 问题类型
	QuestionType int64 `json:"question_type,omitempty" xml:"question_type,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 问题id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 问题类型：0-普通/1-引用
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 问题id
	QuestionId int64 `json:"question_id,omitempty" xml:"question_id,omitempty"`
	// 是否相似标题
	IsSimilarTitle int64 `json:"is_similar_title,omitempty" xml:"is_similar_title,omitempty"`
	// 来源：草稿表-0/正式表-1
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 相似问题主键id
	SimilarId int64 `json:"similar_id,omitempty" xml:"similar_id,omitempty"`
}

var poolKcSearchQuestion = sync.Pool{
	New: func() any {
		return new(KcSearchQuestion)
	},
}

// GetKcSearchQuestion() 从对象池中获取KcSearchQuestion
func GetKcSearchQuestion() *KcSearchQuestion {
	return poolKcSearchQuestion.Get().(*KcSearchQuestion)
}

// ReleaseKcSearchQuestion 释放KcSearchQuestion
func ReleaseKcSearchQuestion(v *KcSearchQuestion) {
	v.CatIdPathList = v.CatIdPathList[:0]
	v.EntityCode = ""
	v.ModifiedUserName = ""
	v.CreateUserName = ""
	v.TitleSegment = ""
	v.Title = ""
	v.GmtModified = ""
	v.GmtCreate = ""
	v.EndDate = ""
	v.StartDate = ""
	v.Ext = ""
	v.CatName = ""
	v.CatPathName = ""
	v.Uuid = ""
	v.Score = ""
	v.SimilarTitle = ""
	v.QuestionType = 0
	v.TenantId = 0
	v.Status = 0
	v.Id = 0
	v.Type = 0
	v.QuestionId = 0
	v.IsSimilarTitle = 0
	v.Source = 0
	v.Version = 0
	v.CatId = 0
	v.SimilarId = 0
	poolKcSearchQuestion.Put(v)
}
