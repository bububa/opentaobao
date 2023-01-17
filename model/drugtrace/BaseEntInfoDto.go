package drugtrace

// BaseEntInfoDto 结构体
type BaseEntInfoDto struct {
	// 上市许可持有人名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 统一社会信用代码（上市许可持有人）
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
}
