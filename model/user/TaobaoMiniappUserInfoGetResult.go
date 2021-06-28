package user

// TaobaoMiniappUserInfoGetResult 
/* model for simplify = false
type TaobaoMiniappUserInfoGetResult struct {

    // model
    
    Model  *struct {
        OpenUserInfoDto  *OpenUserInfoDto `json:"open_user_info_dto,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误信息
    
    ErrMessage   string `json:"err_message,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

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
