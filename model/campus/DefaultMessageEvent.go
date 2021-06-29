package campus

// DefaultMessageEvent 
type DefaultMessageEvent struct {

    // 消息ID(唯一ID)
    
    MsgKey   string `json:"msg_key,omitempty" xml:"msg_key,omitempty"`
    

    // 时间戳
    
    EventTimeMillis   int64 `json:"event_time_millis,omitempty" xml:"event_time_millis,omitempty"`
    

    // 来源
    
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    

    // 接收人用户ID
    
    Receiver   int64 `json:"receiver,omitempty" xml:"receiver,omitempty"`
    

    // 消息事件
    
    Action   string `json:"action,omitempty" xml:"action,omitempty"`
    

    // 消息类型
    
    EventType   string `json:"event_type,omitempty" xml:"event_type,omitempty"`
    

    // 消息参数
    
    JsonData   string `json:"json_data,omitempty" xml:"json_data,omitempty"`
    

    // 资源ID
    
    ResourceId   string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
    

    // 资源URL
    
    ResourceUrl   string `json:"resource_url,omitempty" xml:"resource_url,omitempty"`
    

    // 发送人
    
    Sender   int64 `json:"sender,omitempty" xml:"sender,omitempty"`
    

}
