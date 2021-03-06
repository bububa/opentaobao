package mozivds

// RemoveTenantAdminsRequest 结构体
type RemoveTenantAdminsRequest struct {
	// 人员Code列表
	EmployeeCodes []string `json:"employee_codes,omitempty" xml:"employee_codes>string,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 请求元数据
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 租户Id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}
