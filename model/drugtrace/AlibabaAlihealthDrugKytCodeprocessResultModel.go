package drugtrace

// AlibabaAlihealthDrugKytCodeprocessResultModel 
type AlibabaAlihealthDrugKytCodeprocessResultModel struct {

    // 返回结果
    
    Model   *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
    

    // 错误信息码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 成功标识
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
    

}
