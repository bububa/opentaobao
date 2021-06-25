package user

// TaobaoMiniappUserInfoGetResult 
type TaobaoMiniappUserInfoGetResult struct {

    // model
    Model   *OpenUserInfoDto `json:"model,omitempty"`

    // 错误信息
    ErrMessage   string `json:"err_message,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}