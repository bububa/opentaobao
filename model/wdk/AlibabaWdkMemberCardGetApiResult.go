package wdk

// AlibabaWdkMemberCardGetApiResult 
type AlibabaWdkMemberCardGetApiResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误消息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 错误结果码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 会员信息模型
    
    Model   *MemberInfoDo `json:"model,omitempty" xml:"model,omitempty"`
    

}
