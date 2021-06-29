package uscesl

// TaobaoUsceslBizItemLightUpResult 
type TaobaoUsceslBizItemLightUpResult struct {

    // 返回码
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 业务错误码
    
    BusinessCode   string `json:"business_code,omitempty" xml:"business_code,omitempty"`
    

    // true或false
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    

    // true或false
    
    Target   *LightResultInfoBo `json:"target,omitempty" xml:"target,omitempty"`
    

}
