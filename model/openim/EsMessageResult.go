package openim

// EsMessageResult 
type EsMessageResult struct {

    // 消息序列
    
    Messages   []EsMessage `json:"messages,omitempty" xml:"messages,omitempty"`
    

    // nextkey
    
    NextKey   string `json:"next_key,omitempty" xml:"next_key,omitempty"`
    

}
