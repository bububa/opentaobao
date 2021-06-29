package lsttrade

// HsfResultVo 
type HsfResultVo struct {

    // 返回msg
    
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    

    // 返回状态
    
    ResultStatus   bool `json:"result_status,omitempty" xml:"result_status,omitempty"`
    

    // 返回code
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 内容
    
    Content   bool `json:"content,omitempty" xml:"content,omitempty"`
    

}
