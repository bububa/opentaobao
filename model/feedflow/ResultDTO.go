package feedflow

// ResultDTO 
type ResultDTO struct {

    // 返回说明信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 账户信息
    
    Result   *AccountDTO `json:"result,omitempty" xml:"result,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
