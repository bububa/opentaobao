package security

// RpInitResultBo 结构体
type RpInitResultBo struct {
	// token
	VerifyToken string `json:"verify_token,omitempty" xml:"verify_token,omitempty"`
	// 过期时间
	Expire int64 `json:"expire,omitempty" xml:"expire,omitempty"`
}
