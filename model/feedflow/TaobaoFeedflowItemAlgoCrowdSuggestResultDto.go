package feedflow

// TaobaoFeedflowItemAlgoCrowdSuggestResultDto 
type TaobaoFeedflowItemAlgoCrowdSuggestResultDto struct {

    // 调用是否成功,true-成功，false-失败
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 人群列表
    
    Crowds   []CrowdDto `json:"crowds,omitempty" xml:"crowds,omitempty"`
    

    // 失败时候的消息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

}
