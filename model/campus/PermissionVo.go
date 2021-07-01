package campus

// PermissionVo 结构体
type PermissionVo struct {
	// 权限健
	PrivKey string `json:"priv_key,omitempty" xml:"priv_key,omitempty"`
	// 权限值
	PrivVaue string `json:"priv_vaue,omitempty" xml:"priv_vaue,omitempty"`
	// 权限类型
	PrivType string `json:"priv_type,omitempty" xml:"priv_type,omitempty"`
	// 权限名称
	PrivName string `json:"priv_name,omitempty" xml:"priv_name,omitempty"`
	// 生效开始时间
	EffectiveTime string `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	// 生效过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}
