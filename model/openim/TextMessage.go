package openim

// TextMessage 
/* model for simplify = false
type TextMessage struct {

    // 发送方userid
    
    FromId   string `json:"from_id,omitempty"`
    

    // 消息
    
    Message   string `json:"message,omitempty"`
    

    // 消息时间。UTC时间，精确到秒。必须在一个月内
    
    Time   int64 `json:"time,omitempty"`
    

    // 接受方userid
    
    ToId   string `json:"to_id,omitempty"`
    

}
*/

// TextMessage 
type TextMessage struct {

    // 发送方userid
    FromId   string `json:"from_id,omitempty"`

    // 消息
    Message   string `json:"message,omitempty"`

    // 消息时间。UTC时间，精确到秒。必须在一个月内
    Time   int64 `json:"time,omitempty"`

    // 接受方userid
    ToId   string `json:"to_id,omitempty"`

}
