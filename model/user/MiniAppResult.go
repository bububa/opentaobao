package user

// MiniAppResult 
/* model for simplify = false
type MiniAppResult struct {

    // 错误提示
    
    Msg   string `json:"msg,omitempty"`
    

    // 错误码
    
    Code   string `json:"code,omitempty"`
    

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

    // 用户数据
    
    Data  *struct {
        UserInfoDto  *UserInfoDto `json:"user_info_dto,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

// MiniAppResult 
type MiniAppResult struct {

    // 错误提示
    Msg   string `json:"msg,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

    // 用户数据
    Data   *UserInfoDto `json:"data,omitempty"`

}
