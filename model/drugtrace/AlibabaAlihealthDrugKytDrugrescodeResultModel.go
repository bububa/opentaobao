package drugtrace

// AlibabaAlihealthDrugKytDrugrescodeResultModel 
type AlibabaAlihealthDrugKytDrugrescodeResultModel struct {

    // 返回模型
    
    Model   *Page `json:"model,omitempty" xml:"model,omitempty"`
    

    // 状态码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // 状态值
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 是否响应成功
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
    

}
