package tbtrade

// SecretNo 结构体
type SecretNo struct {
	// 隐私号
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 延期后隐私号的过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}
