package drugtrace

// DrugCode 结构体
type DrugCode struct {
	// 是否有码
	HasCode bool `json:"has_code,omitempty" xml:"has_code,omitempty"`
	// 包装规格
	PkgSpecList []string `json:"pkg_spec_list,omitempty" xml:"pkg_spec_list>string,omitempty"`
}
