package campus

// MobileReceiverMessageEvent 
type MobileReceiverMessageEvent struct {

    // 消息key
    
    MsgKey   string `json:"msg_key,omitempty" xml:"msg_key,omitempty"`
    

    // 时间戳
    
    EventTimeMillis   int64 `json:"event_time_millis,omitempty" xml:"event_time_millis,omitempty"`
    

    // 来源
    
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    

    // 消息事件
    
    Action   string `json:"action,omitempty" xml:"action,omitempty"`
    

    // 接收人手机号
    
    MobileReceiver   string `json:"mobile_receiver,omitempty" xml:"mobile_receiver,omitempty"`
    

    // 消息类型
    
    EventType   string `json:"event_type,omitempty" xml:"event_type,omitempty"`
    

    // 消息内容
    
    JsonData   string `json:"json_data,omitempty" xml:"json_data,omitempty"`
    

    // 资源ID
    
    ResourceId   string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
    

    // 资源URL
    
    ResourceUrl   string `json:"resource_url,omitempty" xml:"resource_url,omitempty"`
    

    // 发送者
    
    Sender   int64 `json:"sender,omitempty" xml:"sender,omitempty"`
    

}
