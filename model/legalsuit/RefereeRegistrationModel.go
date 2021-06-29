package legalsuit

// RefereeRegistrationModel 
type RefereeRegistrationModel struct {
    // 是否开庭
    IsCourtOpened   string `json:"is_court_opened,omitempty" xml:"is_court_opened,omitempty"`
    // 主要观点
    MainPoint   string `json:"main_point,omitempty" xml:"main_point,omitempty"`
    // 撤诉原因(原因)
    TypicalCause   string `json:"typical_cause,omitempty" xml:"typical_cause,omitempty"`
    // 其他业务标签内容
    OtherBusinessLabel   string `json:"other_business_label,omitempty" xml:"other_business_label,omitempty"`
    // 业务标签Keys
    BusinessLabelKeys   string `json:"business_label_keys,omitempty" xml:"business_label_keys,omitempty"`
    // 判决结果二级
    JudgementResultSecond   string `json:"judgement_result_second,omitempty" xml:"judgement_result_second,omitempty"`
    // 判决结果
    JudgementResult   string `json:"judgement_result,omitempty" xml:"judgement_result,omitempty"`
    // 接收时间
    ReceivedTime   string `json:"received_time,omitempty" xml:"received_time,omitempty"`
    // 接口调用时间
    CallingTime   string `json:"calling_time,omitempty" xml:"calling_time,omitempty"`
    // 诉讼系统裁判登记信息
    RegistrationId   int64 `json:"registration_id,omitempty" xml:"registration_id,omitempty"`
    // 案件id
    SuitId   int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
    // 更新人
    Updater   string `json:"updater,omitempty" xml:"updater,omitempty"`
    // 创建人
    Founder   string `json:"founder,omitempty" xml:"founder,omitempty"`
    // 判决金额（其他被告赔付）
    DefendantReceive   string `json:"defendant_receive,omitempty" xml:"defendant_receive,omitempty"`
    // 判决金额（我方获赔）
    AmountOurReceive   string `json:"amount_our_receive,omitempty" xml:"amount_our_receive,omitempty"`
    // 判决金额（我方赔付）
    AmountOurGive   string `json:"amount_our_give,omitempty" xml:"amount_our_give,omitempty"`
    // 诉讼金额
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 我方处理人
    Processor   string `json:"processor,omitempty" xml:"processor,omitempty"`
    // 平台是否需要承担责任
    IsAssumeResponsibility   string `json:"is_assume_responsibility,omitempty" xml:"is_assume_responsibility,omitempty"`
    // 更新时间
    UpdateTime   string `json:"update_time,omitempty" xml:"update_time,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 撤诉日期
    DropDate   string `json:"drop_date,omitempty" xml:"drop_date,omitempty"`
    // 反馈内容
    FeedbackContent   string `json:"feedback_content,omitempty" xml:"feedback_content,omitempty"`
    // 反馈ID
    FeedbackId   int64 `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
    // 反馈附件
    FeedbackAttachmentList   []LFileModel `json:"feedback_attachment_list,omitempty" xml:"feedback_attachment_list>l_file_model,omitempty"`
    // 调解凭证附件列表
    MediationCertificateFiles   []FileModel `json:"mediation_certificate_files,omitempty" xml:"mediation_certificate_files>file_model,omitempty"`
    // 裁定书附件列表
    WrittenVerdictFiles   []LFileModel `json:"written_verdict_files,omitempty" xml:"written_verdict_files>l_file_model,omitempty"`
    // 和解协议附件列表
    SettlementAgreeFiles   []LFileModel `json:"settlement_agree_files,omitempty" xml:"settlement_agree_files>l_file_model,omitempty"`
    // 裁决书附件列表
    JudgeFiles   []LFileModel `json:"judge_files,omitempty" xml:"judge_files>l_file_model,omitempty"`
    // 撤诉文件附件列表
    DroppedFiles   []LFileModel `json:"dropped_files,omitempty" xml:"dropped_files>l_file_model,omitempty"`
    // 其他附件附件列表
    OtherFiles   []LFileModel `json:"other_files,omitempty" xml:"other_files>l_file_model,omitempty"`
    // 委托ID,
    EntrustId   int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
    // 来源平台
    ResultFrom   string `json:"result_from,omitempty" xml:"result_from,omitempty"`
    // 操作类型
    OperationType   string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
}
