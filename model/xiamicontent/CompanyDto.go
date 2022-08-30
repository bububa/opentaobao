package xiamicontent

// CompanyDto 结构体
type CompanyDto struct {
	// 厂牌名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 厂牌id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
}
