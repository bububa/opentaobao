package legalsuit

// DominationDissentModel 
type DominationDissentModel struct {

    // 裁定书附件数量
    
    DominationDissentRulingCount   int64 `json:"domination_dissent_ruling_count,omitempty" xml:"domination_dissent_ruling_count,omitempty"`
    

    // 裁定书附件
    
    DominationDissentRulingList   []FileModel `json:"domination_dissent_ruling_list,omitempty" xml:"domination_dissent_ruling_list,omitempty"`
    

    // 备注
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 管辖异议结果
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    

    // 快递编号
    
    ExpressNum   string `json:"express_num,omitempty" xml:"express_num,omitempty"`
    

    // 快递公司
    
    ExpressCompany   string `json:"express_company,omitempty" xml:"express_company,omitempty"`
    

    // 法官
    
    Judge   string `json:"judge,omitempty" xml:"judge,omitempty"`
    

    // 是否开庭
    
    IsCourt   string `json:"is_court,omitempty" xml:"is_court,omitempty"`
    

    // 反馈附件的个数
    
    FeedbackAttachmentCount   int64 `json:"feedback_attachment_count,omitempty" xml:"feedback_attachment_count,omitempty"`
    

    // 反馈附件
    
    FeedbackAttachmentList   []FileModel `json:"feedback_attachment_list,omitempty" xml:"feedback_attachment_list,omitempty"`
    

    // 反馈内容
    
    FeedbackContent   string `json:"feedback_content,omitempty" xml:"feedback_content,omitempty"`
    

    // 反馈ID
    
    FeedbackId   int64 `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
    

    // 更新时间
    
    UpdateTime   string `json:"update_time,omitempty" xml:"update_time,omitempty"`
    

    // 更新人员
    
    Updater   string `json:"updater,omitempty" xml:"updater,omitempty"`
    

    // 调用时间
    
    CallingTime   string `json:"calling_time,omitempty" xml:"calling_time,omitempty"`
    

    // 案件ID
    
    SuitId   int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
    

    // 委托ID
    
    EntrustId   int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
    

    // 管辖ID
    
    DominationId   int64 `json:"domination_id,omitempty" xml:"domination_id,omitempty"`
    

    // 供应商code
    
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    

    // 管辖文书数量
    
    ApplicationCount   int64 `json:"application_count,omitempty" xml:"application_count,omitempty"`
    

    // 管辖文书
    
    ApplicationAttachmentList   []FileModel `json:"application_attachment_list,omitempty" xml:"application_attachment_list,omitempty"`
    

    // 管辖程序名称
    
    DominationProcess   string `json:"domination_process,omitempty" xml:"domination_process,omitempty"`
    

    // 开庭日期
    
    CourtTime   string `json:"court_time,omitempty" xml:"court_time,omitempty"`
    

    // 操作类型，这里只能更新
    
    OperationType   string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
    

}
