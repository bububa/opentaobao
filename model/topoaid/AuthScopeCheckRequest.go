package topoaid

// AuthScopeCheckRequest 结构体
type AuthScopeCheckRequest struct {
	// 用户手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 从context中获取的用户标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
