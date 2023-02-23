package tmallgenie

// UnionIdInfo 结构体
type UnionIdInfo struct {
	// 组织id
	OrganizationId string `json:"organization_id,omitempty" xml:"organization_id,omitempty"`
	// 开放unionId
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
}
