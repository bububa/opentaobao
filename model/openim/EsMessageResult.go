package openim

// EsMessageResult 
/* model for simplify = false
type EsMessageResult struct {

    // 消息序列
    
    Messages  struct {
        EsMessage  []EsMessage `json:"es_message,omitempty"`
    } `json:"messages,omitempty"`
    

    // nextkey
    
    NextKey   string `json:"next_key,omitempty"`
    

}
*/

// EsMessageResult 
type EsMessageResult struct {

    // 消息序列
    Messages   []EsMessage `json:"messages,omitempty"`

    // nextkey
    NextKey   string `json:"next_key,omitempty"`

}
