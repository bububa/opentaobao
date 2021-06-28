package user

// TokenInfoExt 
/* model for simplify = false
type TokenInfoExt struct {

    // open account当前token info中open account id对应的open account信息
    
    OpenAccount  *struct {
        OpenAccount  *OpenAccount `json:"open_account,omitempty"`
    } `json:"open_account,omitempty"`
    

    // 授权登录后返回的信息
    
    OauthOtherInfo   string `json:"oauth_other_info,omitempty"`
    

}
*/

// TokenInfoExt 
type TokenInfoExt struct {

    // open account当前token info中open account id对应的open account信息
    OpenAccount   *OpenAccount `json:"open_account,omitempty"`

    // 授权登录后返回的信息
    OauthOtherInfo   string `json:"oauth_other_info,omitempty"`

}
