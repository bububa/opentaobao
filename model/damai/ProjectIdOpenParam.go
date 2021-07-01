package damai

// ProjectIdOpenParam 结构体
type ProjectIdOpenParam struct {
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
}
