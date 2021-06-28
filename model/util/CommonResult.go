package util

// CommonResult 
/* model for simplify = false
type CommonResult struct {

    // 服务子错误码
    
    Code   string `json:"code,omitempty"`
    

    // 服务子扩展错误码
    
    SubCode   string `json:"sub_code,omitempty"`
    

    // 成功标识
    
    Success   bool `json:"success,omitempty"`
    

    // 无
    
    Model  *struct {
        AlibabaTaobaoWtUserCrowdModel  *AlibabaTaobaoWtUserCrowdModel `json:"alibaba_taobao_wt_user_crowd_model,omitempty"`
    } `json:"model,omitempty"`
    

    // 描述信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 描述错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 描述请求信息
    
    Desc   string `json:"desc,omitempty"`
    

}
*/

// CommonResult 
type CommonResult struct {

    // 服务子错误码
    Code   string `json:"code,omitempty"`

    // 服务子扩展错误码
    SubCode   string `json:"sub_code,omitempty"`

    // 成功标识
    Success   bool `json:"success,omitempty"`

    // 无
    Model   *AlibabaTaobaoWtUserCrowdModel `json:"model,omitempty"`

    // 描述信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 描述错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 描述请求信息
    Desc   string `json:"desc,omitempty"`

}
