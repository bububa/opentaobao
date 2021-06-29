package tmallgenie

// MessageUnicastRequest 
type MessageUnicastRequest struct {

    // 推送上下文信息
    
    PushContext   string `json:"push_context,omitempty" xml:"push_context,omitempty"`
    

    // 消息发送目标
    
    SendTarget   *MessageSendTarget `json:"send_target,omitempty" xml:"send_target,omitempty"`
    

    // 消息模板id
    
    MessageTemplateId   string `json:"message_template_id,omitempty" xml:"message_template_id,omitempty"`
    

    // 消息号id
    
    MessageEntityId   string `json:"message_entity_id,omitempty" xml:"message_entity_id,omitempty"`
    

    // 模板占位符
    
    PlaceHolder   string `json:"place_holder,omitempty" xml:"place_holder,omitempty"`
    

}
