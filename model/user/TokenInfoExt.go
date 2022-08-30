package user

// TokenInfoExt 结构体
type TokenInfoExt struct {
	// open account当前token info中open account id对应的open account信息
	OpenAccount *OpenAccount `json:"open_account,omitempty" xml:"open_account,omitempty"`
	// oauthOtherInfo
	OauthOtherInfo *OAuthOtherInfo `json:"oauth_other_info,omitempty" xml:"oauth_other_info,omitempty"`
}
