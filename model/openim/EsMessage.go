package openim

// EsMessage 
/* model for simplify = false
type EsMessage struct {

    // 消息时间，UTC时间
    
    Time   int64 `json:"time,omitempty"`
    

    // 消息UUID
    
    Uuid   int64 `json:"uuid,omitempty"`
    

    // 消息类型
    
    Type   int64 `json:"type,omitempty"`
    

    // 发送方
    
    FromId  *struct {
        OpenImUser  *OpenImUser `json:"open_im_user,omitempty"`
    } `json:"from_id,omitempty"`
    

    // 接收方
    
    ToId  *struct {
        OpenImUser  *OpenImUser `json:"open_im_user,omitempty"`
    } `json:"to_id,omitempty"`
    

    // 消息内容
    
    Content  struct {
        RoamingMessageItem  []RoamingMessageItem `json:"roaming_message_item,omitempty"`
    } `json:"content,omitempty"`
    

}
*/

// EsMessage 
type EsMessage struct {

    // 消息时间，UTC时间
    Time   int64 `json:"time,omitempty"`

    // 消息UUID
    Uuid   int64 `json:"uuid,omitempty"`

    // 消息类型
    Type   int64 `json:"type,omitempty"`

    // 发送方
    FromId   *OpenImUser `json:"from_id,omitempty"`

    // 接收方
    ToId   *OpenImUser `json:"to_id,omitempty"`

    // 消息内容
    Content   []RoamingMessageItem `json:"content,omitempty"`

}
