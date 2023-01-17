package topoaid

// AuthScopeCheckResponse 结构体
type AuthScopeCheckResponse struct {
	// 授权列表
	ScopeNames []string `json:"scope_names,omitempty" xml:"scope_names>string,omitempty"`
	// 授权到期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}
