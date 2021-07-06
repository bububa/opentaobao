package moziacl

// BucUserPrincipalParam 结构体
type BucUserPrincipalParam struct {
	// 租户名称，不需要填
	TenantName string `json:"tenant_name,omitempty" xml:"tenant_name,omitempty"`
	// 操作人userId
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 操作人所在租户Id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}
