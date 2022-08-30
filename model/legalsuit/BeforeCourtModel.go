package legalsuit

// BeforeCourtModel 结构体
type BeforeCourtModel struct {
	// 反馈附件信息
	FeedbackAttachmentList []FileModel `json:"feedback_attachment_list,omitempty" xml:"feedback_attachment_list>file_model,omitempty"`
	// 庭前附件信息
	AttachmentList []FileModel `json:"attachment_list,omitempty" xml:"attachment_list>file_model,omitempty"`
	// 庭前沟通
	CommunicateList []CommunicateModel `json:"communicate_list,omitempty" xml:"communicate_list>communicate_model,omitempty"`
	// 诉讼风险
	RiskPredict string `json:"risk_predict,omitempty" xml:"risk_predict,omitempty"`
	// 我方抗辩思路
	MainPoint string `json:"main_point,omitempty" xml:"main_point,omitempty"`
	// 我方观点
	OurQuestions string `json:"our_questions,omitempty" xml:"our_questions,omitempty"`
	// 法庭观点
	CourtQuestions string `json:"court_questions,omitempty" xml:"court_questions,omitempty"`
	// 对方观点
	TheirQuestions string `json:"their_questions,omitempty" xml:"their_questions,omitempty"`
	// 初步考虑对策
	OurAttitude string `json:"our_attitude,omitempty" xml:"our_attitude,omitempty"`
	// 反馈内容
	FeedbackContent string `json:"feedback_content,omitempty" xml:"feedback_content,omitempty"`
	// 更新人
	Updater string `json:"updater,omitempty" xml:"updater,omitempty"`
	// 创建人
	Founder string `json:"founder,omitempty" xml:"founder,omitempty"`
	// 供应商code
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 对方主张
	DefendantEstimate string `json:"defendant_estimate,omitempty" xml:"defendant_estimate,omitempty"`
	// 调用时间
	CallingTime string `json:"calling_time,omitempty" xml:"calling_time,omitempty"`
	// 操作类型
	OperationType string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	// 附件个数
	AttachmentCount int64 `json:"attachment_count,omitempty" xml:"attachment_count,omitempty"`
	// 反馈ID
	FeedbackId int64 `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
	// 庭前ID
	BeforeCourtId int64 `json:"before_court_id,omitempty" xml:"before_court_id,omitempty"`
	// 案件ID
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托ID
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
}
