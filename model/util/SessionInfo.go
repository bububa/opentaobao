package util

// SessionInfo 结构体
type SessionInfo struct {
	// skey
	Skey string `json:"skey,omitempty" xml:"skey,omitempty"`
	// openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// accessToken
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// unionId
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
}
