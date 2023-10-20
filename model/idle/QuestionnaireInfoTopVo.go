package idle

import (
	"sync"
)

// QuestionnaireInfoTopVo 结构体
type QuestionnaireInfoTopVo struct {
	// 投放业务，RECYCLE_3C（回收），RECYCLE_TENDER（寄拍）
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 在线版本号
	OnlineVersion string `json:"online_version,omitempty" xml:"online_version,omitempty"`
	// 在线问卷内容
	OnlineQuestionnaire string `json:"online_questionnaire,omitempty" xml:"online_questionnaire,omitempty"`
	// 预发布版本号
	PreVersion string `json:"pre_version,omitempty" xml:"pre_version,omitempty"`
	// 预发布问卷
	PreQuestionnaire string `json:"pre_questionnaire,omitempty" xml:"pre_questionnaire,omitempty"`
	// 预发模板名称
	PreTemplateName string `json:"pre_template_name,omitempty" xml:"pre_template_name,omitempty"`
	// 线上模板名称
	OnlineTemplateName string `json:"online_template_name,omitempty" xml:"online_template_name,omitempty"`
	// SPUID
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 预发模板ID
	PreTemplateId int64 `json:"pre_template_id,omitempty" xml:"pre_template_id,omitempty"`
	// 线上模板ID
	OnlineTemplateId int64 `json:"online_template_id,omitempty" xml:"online_template_id,omitempty"`
}

var poolQuestionnaireInfoTopVo = sync.Pool{
	New: func() any {
		return new(QuestionnaireInfoTopVo)
	},
}

// GetQuestionnaireInfoTopVo() 从对象池中获取QuestionnaireInfoTopVo
func GetQuestionnaireInfoTopVo() *QuestionnaireInfoTopVo {
	return poolQuestionnaireInfoTopVo.Get().(*QuestionnaireInfoTopVo)
}

// ReleaseQuestionnaireInfoTopVo 释放QuestionnaireInfoTopVo
func ReleaseQuestionnaireInfoTopVo(v *QuestionnaireInfoTopVo) {
	v.BizType = ""
	v.OnlineVersion = ""
	v.OnlineQuestionnaire = ""
	v.PreVersion = ""
	v.PreQuestionnaire = ""
	v.PreTemplateName = ""
	v.OnlineTemplateName = ""
	v.SpuId = 0
	v.PreTemplateId = 0
	v.OnlineTemplateId = 0
	poolQuestionnaireInfoTopVo.Put(v)
}
