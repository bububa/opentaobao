package cainiaohandover

// UserInfoDto 
type UserInfoDto struct {
    // 每个商家在ISV系统的唯一标识，一般为商家ISV账号的id
    TopUserKey   string `json:"top_user_key,omitempty" xml:"top_user_key,omitempty"`
}
