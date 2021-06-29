package alsc

// TemplateNotifyRequest 
type TemplateNotifyRequest struct {

    // 用户OpenId
    
    OpenId   string `json:"open_id,omitempty" xml:"open_id,omitempty"`
    

    // 小程序appid
    
    Appid   string `json:"appid,omitempty" xml:"appid,omitempty"`
    

    // 消息主体码
    
    NotifyId   string `json:"notify_id,omitempty" xml:"notify_id,omitempty"`
    

    // 业务标签
    
    BizTag   string `json:"biz_tag,omitempty" xml:"biz_tag,omitempty"`
    

    // 模板变量数据
    
    TemplateArgs   string `json:"template_args,omitempty" xml:"template_args,omitempty"`
    

}
