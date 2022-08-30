package idle

// QuestionnaireInfoTopVO 结构体
type QuestionnaireInfoTopVO struct {
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
	// SPUID
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}
