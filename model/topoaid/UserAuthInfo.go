package topoaid

// UserAuthInfo 结构体
type UserAuthInfo struct {
	// 收件人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 授权信息过期时间
	AuthorizeExpireTime string `json:"authorize_expire_time,omitempty" xml:"authorize_expire_time,omitempty"`
	// 用户唯一标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 虚拟号
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 虚拟号过期时间
	SecretExpireTime string `json:"secret_expire_time,omitempty" xml:"secret_expire_time,omitempty"`
}
