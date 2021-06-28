package openim

// TribeMessage 
type TribeMessage struct {

    // 消息内容节点序列
    
    Content   []MessageItem `json:"content,omitempty" xml:"content,omitempty"`
    

    // 发送方
    
    FromId   *User `json:"from_id,omitempty" xml:"from_id,omitempty"`
    

    // 消息时间。UTC时间
    
    Time   int64 `json:"time,omitempty" xml:"time,omitempty"`
    

    // 消息类型
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 消息UUID
    
    Uuid   int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
    

}
