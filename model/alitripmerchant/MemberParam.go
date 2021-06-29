package alitripmerchant

// MemberParam 
type MemberParam struct {

    // 用户英文姓
    
    LastName   string `json:"last_name,omitempty" xml:"last_name,omitempty"`
    

    // 用户英文名
    
    FirstName   string `json:"first_name,omitempty" xml:"first_name,omitempty"`
    

    // 称呼
    
    Civility   string `json:"civility,omitempty" xml:"civility,omitempty"`
    

    // 用户手机号前缀
    
    PhonePre   string `json:"phone_pre,omitempty" xml:"phone_pre,omitempty"`
    

    // 用户语言
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

    // 是否接受协议
    
    AcceptedTandC   bool `json:"accepted_tand_c,omitempty" xml:"accepted_tand_c,omitempty"`
    

    // 用户手机号
    
    PhoneNum   string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
    

    // 是否接受消息推送
    
    Subscription   bool `json:"subscription,omitempty" xml:"subscription,omitempty"`
    

    // 用户邮箱
    
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    

}
