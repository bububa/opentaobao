package legalsuit

// CaseModel 结构体
type CaseModel struct {
	// 原告信息
	Accusers []AccuserModel `json:"accusers,omitempty" xml:"accusers>accuser_model,omitempty"`
	// 原告诉请
	AccuserAppealList []string `json:"accuser_appeal_list,omitempty" xml:"accuser_appeal_list>string,omitempty"`
	// 被告
	DefendantList []AppelleeModel `json:"defendant_list,omitempty" xml:"defendant_list>appellee_model,omitempty"`
	// 第三人
	ThirdList []LitigantThirdPartyModel `json:"third_list,omitempty" xml:"third_list>litigant_third_party_model,omitempty"`
	// 案件事实调查附件
	CaseFactFileList []FileModel `json:"case_fact_file_list,omitempty" xml:"case_fact_file_list>file_model,omitempty"`
	// 其他附件
	OtherFileList []FileModel `json:"other_file_list,omitempty" xml:"other_file_list>file_model,omitempty"`
	// 证据附件
	EvidenceFileList []FileModel `json:"evidence_file_list,omitempty" xml:"evidence_file_list>file_model,omitempty"`
	// 法院相关附件
	CourtFileList []FileModel `json:"court_file_list,omitempty" xml:"court_file_list>file_model,omitempty"`
	// 起诉状附件
	PleadingFileList []FileModel `json:"pleading_file_list,omitempty" xml:"pleading_file_list>file_model,omitempty"`
	// 预立案号
	PreCaseNumber string `json:"pre_case_number,omitempty" xml:"pre_case_number,omitempty"`
	// 标签4
	Tag4 string `json:"tag4,omitempty" xml:"tag4,omitempty"`
	// 标签3
	Tag3 string `json:"tag3,omitempty" xml:"tag3,omitempty"`
	// 标签2
	Tag2 string `json:"tag2,omitempty" xml:"tag2,omitempty"`
	// 标签1
	Tag1 string `json:"tag1,omitempty" xml:"tag1,omitempty"`
	// 原告主张的法律依据
	AccuserClaimLegalBasis string `json:"accuser_claim_legal_basis,omitempty" xml:"accuser_claim_legal_basis,omitempty"`
	// 原告主张的事实
	AccuserClaimFact string `json:"accuser_claim_fact,omitempty" xml:"accuser_claim_fact,omitempty"`
	// 诉讼请求
	SuitRequest string `json:"suit_request,omitempty" xml:"suit_request,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 当前跟进人
	FollowupPeopleName string `json:"followup_people_name,omitempty" xml:"followup_people_name,omitempty"`
	// 案件归属OU
	OuName string `json:"ou_name,omitempty" xml:"ou_name,omitempty"`
	// 开庭时间
	CourtTime string `json:"court_time,omitempty" xml:"court_time,omitempty"`
	// 起诉时间
	SueTime string `json:"sue_time,omitempty" xml:"sue_time,omitempty"`
	// 案件的货币类型
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 标的金额
	CaseAmount string `json:"case_amount,omitempty" xml:"case_amount,omitempty"`
	// 审判类型
	AuditType string `json:"audit_type,omitempty" xml:"audit_type,omitempty"`
	// 案由
	CaseCause string `json:"case_cause,omitempty" xml:"case_cause,omitempty"`
	// 案件类型2
	CaseType2 string `json:"case_type2,omitempty" xml:"case_type2,omitempty"`
	// 案件类型1
	CaseType1 string `json:"case_type1,omitempty" xml:"case_type1,omitempty"`
	// 诉讼类型
	SuitType string `json:"suit_type,omitempty" xml:"suit_type,omitempty"`
	// 案情描述
	CaseDetailDescription string `json:"case_detail_description,omitempty" xml:"case_detail_description,omitempty"`
	// 案情编号
	CaseDetailCode string `json:"case_detail_code,omitempty" xml:"case_detail_code,omitempty"`
	// 案号
	CaseNumber string `json:"case_number,omitempty" xml:"case_number,omitempty"`
	// 费用信息
	SuitFee *SuitFeeModel `json:"suit_fee,omitempty" xml:"suit_fee,omitempty"`
	// 收案信息
	CheckInMsg *CheckInModel `json:"check_in_msg,omitempty" xml:"check_in_msg,omitempty"`
	// 法院信息
	CourtModel *CourtModel `json:"court_model,omitempty" xml:"court_model,omitempty"`
	// 案件id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
