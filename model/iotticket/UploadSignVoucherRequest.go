package iotticket

// UploadSignVoucherRequest 
type UploadSignVoucherRequest struct {

    // 扩展信息
    
    Feature   string `json:"feature,omitempty" xml:"feature,omitempty"`
    

    // 操作人联系方式
    
    OperatorPhone   string `json:"operator_phone,omitempty" xml:"operator_phone,omitempty"`
    

    // 操作人编码
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

    // 操作人名称
    
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    

    // 服务商唯一编码
    
    SpCode   string `json:"sp_code,omitempty" xml:"sp_code,omitempty"`
    

    // 工单Id
    
    TicketId   int64 `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
    

    // 备注
    
    Comment   string `json:"comment,omitempty" xml:"comment,omitempty"`
    

    // 邮件编码
    
    MailNo   string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
    

    // 快递签收凭证
    
    SignProofs   []string `json:"sign_proofs,omitempty" xml:"sign_proofs>string,omitempty"`
    

}
