package cmns

// MessageAckResult 
/* model for simplify = false
type MessageAckResult struct {

    // 消息回复时间
    
    AckTime   string `json:"ack_time,omitempty"`
    

    // 设备id
    
    Did   int64 `json:"did,omitempty"`
    

    // 消息id
    
    Mid   int64 `json:"mid,omitempty"`
    

    // uuid
    
    Uuid   string `json:"uuid,omitempty"`
    

}
*/

// MessageAckResult 
type MessageAckResult struct {

    // 消息回复时间
    AckTime   string `json:"ack_time,omitempty"`

    // 设备id
    Did   int64 `json:"did,omitempty"`

    // 消息id
    Mid   int64 `json:"mid,omitempty"`

    // uuid
    Uuid   string `json:"uuid,omitempty"`

}
