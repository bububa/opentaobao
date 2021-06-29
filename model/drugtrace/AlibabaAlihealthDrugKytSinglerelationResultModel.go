package drugtrace

// AlibabaAlihealthDrugKytSinglerelationResultModel 
type AlibabaAlihealthDrugKytSinglerelationResultModel struct {

    // model
    
    ModelList   []CodeRelationDto `json:"model_list,omitempty" xml:"model_list,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 是否成功
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
    

}
