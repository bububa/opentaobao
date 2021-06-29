package drugtrace

// Childcodes 
type Childcodes struct {
    // 子码
    PkgLevel   int64 `json:"pkg_level,omitempty" xml:"pkg_level,omitempty"`
    // 子码级别
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
