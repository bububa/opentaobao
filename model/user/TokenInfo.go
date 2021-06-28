package user

// TokenInfo 
/* model for simplify = false
type TokenInfo struct {

    // 时间戳
    
    Timestamp   int64 `json:"timestamp,omitempty"`
    

    // open account id
    
    OpenAccountId   int64 `json:"open_account_id,omitempty"`
    

    // isv自己账号的唯一id
    
    IsvAccountId   string `json:"isv_account_id,omitempty"`
    

    // 用于防重放的唯一id
    
    Uuid   string `json:"uuid,omitempty"`
    

    // ISV APP的登录态时长
    
    LoginStateExpireIn   int64 `json:"login_state_expire_in,omitempty"`
    

    // token info扩展信息
    
    Ext  *struct {
        TokenInfoExt  *TokenInfoExt `json:"token_info_ext,omitempty"`
    } `json:"ext,omitempty"`
    

}
*/

// TokenInfo 
type TokenInfo struct {

    // 时间戳
    Timestamp   int64 `json:"timestamp,omitempty"`

    // open account id
    OpenAccountId   int64 `json:"open_account_id,omitempty"`

    // isv自己账号的唯一id
    IsvAccountId   string `json:"isv_account_id,omitempty"`

    // 用于防重放的唯一id
    Uuid   string `json:"uuid,omitempty"`

    // ISV APP的登录态时长
    LoginStateExpireIn   int64 `json:"login_state_expire_in,omitempty"`

    // token info扩展信息
    Ext   *TokenInfoExt `json:"ext,omitempty"`

}
