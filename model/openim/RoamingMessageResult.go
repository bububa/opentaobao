package openim

// RoamingMessageResult 
/* model for simplify = false
type RoamingMessageResult struct {

    // 下次迭代key
    
    NextKey   string `json:"next_key,omitempty"`
    

    // 消息列表
    
    Messages  struct {
        RoamingMessage  []RoamingMessage `json:"roaming_message,omitempty"`
    } `json:"messages,omitempty"`
    

}
*/

// RoamingMessageResult 
type RoamingMessageResult struct {

    // 下次迭代key
    NextKey   string `json:"next_key,omitempty"`

    // 消息列表
    Messages   []RoamingMessage `json:"messages,omitempty"`

}
