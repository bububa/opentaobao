package user

// TokenInfoExt 
type TokenInfoExt struct {
    // open account当前token info中open account id对应的open account信息
    OpenAccount   *OpenAccount `json:"open_account,omitempty" xml:"open_account,omitempty"`
    // 授权登录后返回的信息
    OauthOtherInfo   string `json:"oauth_other_info,omitempty" xml:"oauth_other_info,omitempty"`
}
