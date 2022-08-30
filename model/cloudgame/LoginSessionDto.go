package cloudgame

// LoginSessionDto 结构体
type LoginSessionDto struct {
	// user的acesstoken
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// user在mp上的accountid
	MpAccountId int64 `json:"mp_account_id,omitempty" xml:"mp_account_id,omitempty"`
}
