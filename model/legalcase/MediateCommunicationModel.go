package legalcase

// MediateCommunicationModel 结构体
type MediateCommunicationModel struct {
	// 附件
	AttachmentList []FileModel `json:"attachment_list,omitempty" xml:"attachment_list>file_model,omitempty"`
	// 沟通记录
	CommunicateRecord string `json:"communicate_record,omitempty" xml:"communicate_record,omitempty"`
	// 联系电话
	ContactNumber string `json:"contact_number,omitempty" xml:"contact_number,omitempty"`
	// 调解联系人
	ContactPeople string `json:"contact_people,omitempty" xml:"contact_people,omitempty"`
	// 调解联系时间
	ContactTime string `json:"contact_time,omitempty" xml:"contact_time,omitempty"`
	// id新增不用
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 调解金额
	MediateAmount string `json:"mediate_amount,omitempty" xml:"mediate_amount,omitempty"`
	// 调解阶段
	Phase string `json:"phase,omitempty" xml:"phase,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 结果原因
	ResultReason string `json:"result_reason,omitempty" xml:"result_reason,omitempty"`
	// 卖家是否要求积极应诉
	SellerAskRespondent bool `json:"seller_ask_respondent,omitempty" xml:"seller_ask_respondent,omitempty"`
	// 解决方案
	Solution string `json:"solution,omitempty" xml:"solution,omitempty"`
}
