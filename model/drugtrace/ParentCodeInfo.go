package drugtrace

// ParentCodeInfo 结构体
type ParentCodeInfo struct {
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 码级别
	PackageLevel int64 `json:"package_level,omitempty" xml:"package_level,omitempty"`
}
