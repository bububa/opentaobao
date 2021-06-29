package drugtrace

// AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel 
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel struct {

    // 对象模型信息
    
    Model   *BillInOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
    

    // 消息码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // 消息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 成功失败
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
    

}
