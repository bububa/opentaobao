package mozi

// ListAccountsByAccountIdsRequest 结构体
type ListAccountsByAccountIdsRequest struct {
	// 账号ID列表
	AccountIds []int64 `json:"account_ids,omitempty" xml:"account_ids>int64,omitempty"`
	// 账号是否可用
	Available string `json:"available,omitempty" xml:"available,omitempty"`
	// 附加信息
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}
