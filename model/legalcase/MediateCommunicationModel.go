package legalcase

// MediateCommunicationModel 
/* model for simplify = false
type MediateCommunicationModel struct {

    // 附件
    
    AttachmentList  struct {
        FileModel  []FileModel `json:"file_model,omitempty"`
    } `json:"attachment_list,omitempty"`
    

    // 沟通记录
    
    CommunicateRecord   string `json:"communicate_record,omitempty"`
    

    // 联系电话
    
    ContactNumber   string `json:"contact_number,omitempty"`
    

    // 调解联系人
    
    ContactPeople   string `json:"contact_people,omitempty"`
    

    // 调解联系时间
    
    ContactTime   string `json:"contact_time,omitempty"`
    

    // id新增不用
    
    Id   int64 `json:"id,omitempty"`
    

    // 调解金额
    
    MediateAmount   string `json:"mediate_amount,omitempty"`
    

    // 调解阶段
    
    Phase   string `json:"phase,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 结果
    
    Result   string `json:"result,omitempty"`
    

    // 结果原因
    
    ResultReason   string `json:"result_reason,omitempty"`
    

    // 卖家是否要求积极应诉
    
    SellerAskRespondent   bool `json:"seller_ask_respondent,omitempty"`
    

    // 解决方案
    
    Solution   string `json:"solution,omitempty"`
    

}
*/

// MediateCommunicationModel 
type MediateCommunicationModel struct {

    // 附件
    AttachmentList   []FileModel `json:"attachment_list,omitempty"`

    // 沟通记录
    CommunicateRecord   string `json:"communicate_record,omitempty"`

    // 联系电话
    ContactNumber   string `json:"contact_number,omitempty"`

    // 调解联系人
    ContactPeople   string `json:"contact_people,omitempty"`

    // 调解联系时间
    ContactTime   string `json:"contact_time,omitempty"`

    // id新增不用
    Id   int64 `json:"id,omitempty"`

    // 调解金额
    MediateAmount   string `json:"mediate_amount,omitempty"`

    // 调解阶段
    Phase   string `json:"phase,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 结果
    Result   string `json:"result,omitempty"`

    // 结果原因
    ResultReason   string `json:"result_reason,omitempty"`

    // 卖家是否要求积极应诉
    SellerAskRespondent   bool `json:"seller_ask_respondent,omitempty"`

    // 解决方案
    Solution   string `json:"solution,omitempty"`

}
