package alicom

// SmsInterceptRequest 
type SmsInterceptRequest struct {

    // 中间号
    
    SecretNo   string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
    

    // 短信发送主叫号码
    
    CallNo   string `json:"call_no,omitempty" xml:"call_no,omitempty"`
    

    // 短信内容，请使用UCS2进行编码
    
    SmsContent   string `json:"sms_content,omitempty" xml:"sms_content,omitempty"`
    

    // 短信时间撮
    
    MtTime   string `json:"mt_time,omitempty" xml:"mt_time,omitempty"`
    

    // 每次呼叫行为和短信行为的唯一ID
    
    CallId   string `json:"call_id,omitempty" xml:"call_id,omitempty"`
    

    // 对应阿里侧的绑定关系ID
    
    SubsId   string `json:"subs_id,omitempty" xml:"subs_id,omitempty"`
    

    // 分配给供应商的KEY
    
    VendorKey   string `json:"vendor_key,omitempty" xml:"vendor_key,omitempty"`
    

}
