package alicom

// OssTokenGetResult 
type OssTokenGetResult struct {
    // model
    Model   *OssTokenResponse `json:"model,omitempty" xml:"model,omitempty"`
    // desc
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
