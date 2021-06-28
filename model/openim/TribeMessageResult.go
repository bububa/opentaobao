package openim

// TribeMessageResult 
/* model for simplify = false
type TribeMessageResult struct {

    // 消息列表
    
    Messages  struct {
        TribeMessage  []TribeMessage `json:"tribe_message,omitempty"`
    } `json:"messages,omitempty"`
    

    // 迭代key
    
    NextKey   string `json:"next_key,omitempty"`
    

}
*/

// TribeMessageResult 
type TribeMessageResult struct {

    // 消息列表
    Messages   []TribeMessage `json:"messages,omitempty"`

    // 迭代key
    NextKey   string `json:"next_key,omitempty"`

}
