package legalsuit

// CourtInfoModel 结构体
type CourtInfoModel struct {
	// 总结/教训
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 合议庭
	FullCourt string `json:"full_court,omitempty" xml:"full_court,omitempty"`
	// 反馈附件
	FeedbackAttachmentList []FileModel `json:"feedback_attachment_list,omitempty" xml:"feedback_attachment_list>file_model,omitempty"`
	// 开庭ID
	CourtId int64 `json:"court_id,omitempty" xml:"court_id,omitempty"`
	// 情况说明
	Show string `json:"show,omitempty" xml:"show,omitempty"`
	// 法官认职
	JudgePosition string `json:"judge_position,omitempty" xml:"judge_position,omitempty"`
	// 法院提问的有价值问题(主诉)
	CourtQuestionAndAnswer string `json:"court_question_and_answer,omitempty" xml:"court_question_and_answer,omitempty"`
	// 案件ID
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 是否网上开庭
	IsOnline string `json:"is_online,omitempty" xml:"is_online,omitempty"`
	// 反馈内容
	FeedbackContent string `json:"feedback_content,omitempty" xml:"feedback_content,omitempty"`
	// 争议焦点归纳
	DisputeSummary string `json:"dispute_summary,omitempty" xml:"dispute_summary,omitempty"`
	// 开庭时间
	CourtTime string `json:"court_time,omitempty" xml:"court_time,omitempty"`
	// 主审法官
	Judge string `json:"judge,omitempty" xml:"judge,omitempty"`
	// 委托ID
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 创建人
	Founder string `json:"founder,omitempty" xml:"founder,omitempty"`
	// 更新人
	Updater string `json:"updater,omitempty" xml:"updater,omitempty"`
	// 起诉阿里的原因
	SueAliReason string `json:"sue_ali_reason,omitempty" xml:"sue_ali_reason,omitempty"`
	// 是否知道投诉渠道
	IsKnowComplainWay string `json:"is_know_complain_way,omitempty" xml:"is_know_complain_way,omitempty"`
	// 非正式开庭原因
	InformalReason string `json:"informal_reason,omitempty" xml:"informal_reason,omitempty"`
	// 问题列表
	QuestionList []CourtProblemModel `json:"question_list,omitempty" xml:"question_list>court_problem_model,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 非正式开庭附件
	FormalAttachmentList []FileModel `json:"formal_attachment_list,omitempty" xml:"formal_attachment_list>file_model,omitempty"`
	// 开庭证据列表
	InclusionOriDefList []CourtEvidenceModel `json:"inclusion_ori_def_list,omitempty" xml:"inclusion_ori_def_list>court_evidence_model,omitempty"`
	// 风险反馈
	RiskFeedback string `json:"risk_feedback,omitempty" xml:"risk_feedback,omitempty"`
	// 非正式开庭附件数量
	InformalAttachmentCount int64 `json:"informal_attachment_count,omitempty" xml:"informal_attachment_count,omitempty"`
	// 案由
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 庭后计划
	Plan string `json:"plan,omitempty" xml:"plan,omitempty"`
	// 开庭方式
	CourtWay string `json:"court_way,omitempty" xml:"court_way,omitempty"`
	// 操作类型
	OperationType string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	// 原告信息
	AccuserList []CourtPartyModel `json:"accuser_list,omitempty" xml:"accuser_list>court_party_model,omitempty"`
	// 原被告是否到庭(主诉)
	AccuserDefendantIsCourt string `json:"accuser_defendant_is_court,omitempty" xml:"accuser_defendant_is_court,omitempty"`
	// 非正式开庭附件
	InformalAttachmentList []FileModel `json:"informal_attachment_list,omitempty" xml:"informal_attachment_list>file_model,omitempty"`
	// 调用时间
	CallingTime string `json:"calling_time,omitempty" xml:"calling_time,omitempty"`
	// 是否有过投诉
	IsComplain string `json:"is_complain,omitempty" xml:"is_complain,omitempty"`
	// 其他特殊情况
	OtherSpecialPosition string `json:"other_special_position,omitempty" xml:"other_special_position,omitempty"`
	// 正式开庭附件数量
	FormalAttachmentCount int64 `json:"formal_attachment_count,omitempty" xml:"formal_attachment_count,omitempty"`
	// 行政区划
	CourtGeographyName string `json:"court_geography_name,omitempty" xml:"court_geography_name,omitempty"`
	// 反馈ID
	FeedbackId int64 `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
	// 被告信息
	DefendantList []CourtPartyModel `json:"defendant_list,omitempty" xml:"defendant_list>court_party_model,omitempty"`
}
