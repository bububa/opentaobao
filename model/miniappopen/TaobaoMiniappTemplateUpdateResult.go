package miniappopen

// TaobaoMiniappTemplateUpdateResult 
type TaobaoMiniappTemplateUpdateResult struct {

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    

    // model
    
    Model   *MiniAppEntityTemplateDto `json:"model,omitempty" xml:"model,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
