package legalsuit

import (
	"sync"
)

// AfterCourtInfoModel 结构体
type AfterCourtInfoModel struct {
	// 附件列表
	AttachmentList []FileModel `json:"attachment_list,omitempty" xml:"attachment_list>file_model,omitempty"`
	// 与当事人沟通内容
	PartyList []FileModel `json:"party_list,omitempty" xml:"party_list>file_model,omitempty"`
	// 与法院沟通内容
	CourtsList []FileModel `json:"courts_list,omitempty" xml:"courts_list>file_model,omitempty"`
	// 非证据材料
	NonEvidenceList []AfterCourtDetailModel `json:"non_evidence_list,omitempty" xml:"non_evidence_list>after_court_detail_model,omitempty"`
	// 证据材料
	EvidenceList []AfterCourtDetailModel `json:"evidence_list,omitempty" xml:"evidence_list>after_court_detail_model,omitempty"`
	// 反馈附件
	FeedbackAttachmentList []FileModel `json:"feedback_attachment_list,omitempty" xml:"feedback_attachment_list>file_model,omitempty"`
	// 预测结果
	PredictionResult string `json:"prediction_result,omitempty" xml:"prediction_result,omitempty"`
	// 交流时间
	CommunicationDate string `json:"communication_date,omitempty" xml:"communication_date,omitempty"`
	// 反馈内容
	FeedbackContent string `json:"feedback_content,omitempty" xml:"feedback_content,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 更新人
	Updater string `json:"updater,omitempty" xml:"updater,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 创建人
	Founder string `json:"founder,omitempty" xml:"founder,omitempty"`
	// 沟通总结
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 代理词要点
	AgentWord string `json:"agent_word,omitempty" xml:"agent_word,omitempty"`
	// 调用时间
	CallingTime string `json:"calling_time,omitempty" xml:"calling_time,omitempty"`
	// 操作类型
	OperationType string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	// 备注（主诉）
	Discription string `json:"discription,omitempty" xml:"discription,omitempty"`
	// 附件数量
	AttachmentCount int64 `json:"attachment_count,omitempty" xml:"attachment_count,omitempty"`
	// 附件数量
	FeedbackAttachmentCount int64 `json:"feedback_attachment_count,omitempty" xml:"feedback_attachment_count,omitempty"`
	// 反馈id
	FeedbackId int64 `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
	// 案件id
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 庭后信息ID,更新时不能为空
	AfterCourtId int64 `json:"after_court_id,omitempty" xml:"after_court_id,omitempty"`
}

var poolAfterCourtInfoModel = sync.Pool{
	New: func() any {
		return new(AfterCourtInfoModel)
	},
}

// GetAfterCourtInfoModel() 从对象池中获取AfterCourtInfoModel
func GetAfterCourtInfoModel() *AfterCourtInfoModel {
	return poolAfterCourtInfoModel.Get().(*AfterCourtInfoModel)
}

// ReleaseAfterCourtInfoModel 释放AfterCourtInfoModel
func ReleaseAfterCourtInfoModel(v *AfterCourtInfoModel) {
	v.AttachmentList = v.AttachmentList[:0]
	v.PartyList = v.PartyList[:0]
	v.CourtsList = v.CourtsList[:0]
	v.NonEvidenceList = v.NonEvidenceList[:0]
	v.EvidenceList = v.EvidenceList[:0]
	v.FeedbackAttachmentList = v.FeedbackAttachmentList[:0]
	v.PredictionResult = ""
	v.CommunicationDate = ""
	v.FeedbackContent = ""
	v.UpdateTime = ""
	v.Updater = ""
	v.CreateTime = ""
	v.Founder = ""
	v.Summary = ""
	v.AgentWord = ""
	v.CallingTime = ""
	v.OperationType = ""
	v.Discription = ""
	v.AttachmentCount = 0
	v.FeedbackAttachmentCount = 0
	v.FeedbackId = 0
	v.SuitId = 0
	v.EntrustId = 0
	v.AfterCourtId = 0
	poolAfterCourtInfoModel.Put(v)
}
