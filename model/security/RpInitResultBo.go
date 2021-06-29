package security

// RpInitResultBo 
type RpInitResultBo struct {
    // 过期时间
    Expire   int64 `json:"expire,omitempty" xml:"expire,omitempty"`
    // token
    VerifyToken   string `json:"verify_token,omitempty" xml:"verify_token,omitempty"`
}
