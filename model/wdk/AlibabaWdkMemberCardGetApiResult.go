package wdk

// AlibabaWdkMemberCardGetApiResult 
/* model for simplify = false
type AlibabaWdkMemberCardGetApiResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误消息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 错误结果码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 会员信息模型
    
    Model  *struct {
        MemberInfoDo  *MemberInfoDo `json:"member_info_do,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

// AlibabaWdkMemberCardGetApiResult 
type AlibabaWdkMemberCardGetApiResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误消息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 错误结果码
    ErrCode   string `json:"err_code,omitempty"`

    // 会员信息模型
    Model   *MemberInfoDo `json:"model,omitempty"`

}
