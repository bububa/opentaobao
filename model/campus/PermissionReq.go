package campus

// PermissionReq 结构体
type PermissionReq struct {
	// 权限key
	PrivKey string `json:"priv_key,omitempty" xml:"priv_key,omitempty"`
	// 生效时间
	EffectiveTime string `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	// 过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}
