package drugtrace

// Childcodes 结构体
type Childcodes struct {
	// 子码级别
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 子码
	PkgLevel int64 `json:"pkg_level,omitempty" xml:"pkg_level,omitempty"`
}
