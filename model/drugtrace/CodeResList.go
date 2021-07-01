package drugtrace

// CodeResList 结构体
type CodeResList struct {
	// 码前缀
	CodePrefix string `json:"code_prefix,omitempty" xml:"code_prefix,omitempty"`
	// 资源码
	ResCode string `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 层级
	CodeLevel string `json:"code_level,omitempty" xml:"code_level,omitempty"`
	// 包装比例
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
}
