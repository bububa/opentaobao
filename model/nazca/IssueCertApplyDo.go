package nazca

// IssueCertApplyDo 
/* model for simplify = false
type IssueCertApplyDo struct {

    // 合同编号
    
    ContractNum   string `json:"contract_num,omitempty"`
    

    // pageNum
    
    PageNum   int64 `json:"page_num,omitempty"`
    

    // 客户在1688的唯一标识
    
    PlatformUserId   string `json:"platform_user_id,omitempty"`
    

    // 页面跳转同步通知URL
    
    ReturnUrl   string `json:"return_url,omitempty"`
    

    // 角色 0：接收者 1：发送者
    
    SendReceive   int64 `json:"send_receive,omitempty"`
    

    // 合同名称
    
    Topic   string `json:"topic,omitempty"`
    

}
*/

// IssueCertApplyDo 
type IssueCertApplyDo struct {

    // 合同编号
    ContractNum   string `json:"contract_num,omitempty"`

    // pageNum
    PageNum   int64 `json:"page_num,omitempty"`

    // 客户在1688的唯一标识
    PlatformUserId   string `json:"platform_user_id,omitempty"`

    // 页面跳转同步通知URL
    ReturnUrl   string `json:"return_url,omitempty"`

    // 角色 0：接收者 1：发送者
    SendReceive   int64 `json:"send_receive,omitempty"`

    // 合同名称
    Topic   string `json:"topic,omitempty"`

}
