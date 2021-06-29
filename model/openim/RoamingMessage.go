package openim

// RoamingMessage 
type RoamingMessage struct {
    // 消息内容
    ContentList   []RoamingMessageItem `json:"content_list,omitempty" xml:"content_list>roaming_message_item,omitempty"`
    // 消息时间（UTC时间）
    Time   int64 `json:"time,omitempty" xml:"time,omitempty"`
    // 消息方向。user1 -> user2 = 0 , user2->user1 = 1
    Direction   int64 `json:"direction,omitempty" xml:"direction,omitempty"`
    // 消息唯一ID
    Uuid   int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
    // 消息类型
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
}
