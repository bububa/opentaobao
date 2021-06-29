package moziacl

// BucPrincipalParam 
type BucPrincipalParam struct {
    // 操作人所在租户
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 操作人userId
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}
