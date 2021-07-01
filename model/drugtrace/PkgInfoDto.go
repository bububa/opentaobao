package drugtrace

// PkgInfoDto 结构体
type PkgInfoDto struct {
	// 码信息
	CodeList []string `json:"code_list,omitempty" xml:"code_list>string,omitempty"`
}
