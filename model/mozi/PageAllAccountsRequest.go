package mozi

// PageAllAccountsRequest 结构体
type PageAllAccountsRequest struct {
	// 账号类型，21 钉钉 11 淘宝
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 是否可用 T可用 F不可用，不填是全部
	Available string `json:"available,omitempty" xml:"available,omitempty"`
	// 附加信息
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 是否在职 A在职 I离职，不填是全部
	HrStatus string `json:"hr_status,omitempty" xml:"hr_status,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 每页的大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页号
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 0表示未激活，1表示激活，不填是全部
	ActiveLevel int64 `json:"active_level,omitempty" xml:"active_level,omitempty"`
	// 状态，0表示正常，5是离职回收，104是冻结，不填是全部
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否返回总数
	ReturnTotalSize bool `json:"return_total_size,omitempty" xml:"return_total_size,omitempty"`
}
