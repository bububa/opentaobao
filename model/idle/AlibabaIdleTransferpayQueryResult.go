package idle

// AlibabaIdleTransferpayQueryResult 
type AlibabaIdleTransferpayQueryResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 详情数据
    
    Module   *Serializable `json:"module,omitempty" xml:"module,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

}
