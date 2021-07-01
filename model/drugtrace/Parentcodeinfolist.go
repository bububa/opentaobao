package drugtrace

// Parentcodeinfolist 结构体
type Parentcodeinfolist struct {
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 父码级别
	PackageLevel string `json:"package_level,omitempty" xml:"package_level,omitempty"`
}
